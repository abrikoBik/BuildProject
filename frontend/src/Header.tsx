import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import axios from "axios";

const Header = () => {

    const [userRole, setUserRole] = useState<string>();

    const [userBlock, setUserBlock] = useState(false);
    const [userBlockVisible, setUseBlockVisible] = useState(false);

    const [name, setName] = useState<string>("");

    const getUserRole = async () => {
        try {
            const roleRequest = await axios.get(
                `/api/users/${localStorage.getItem("userID")}/worker`
            );

            setUserRole(roleRequest.data.role);
        } catch (e : unknown) {
            console.log(e);
            setUserRole("client");
        }
    } 

    useEffect(() => {
        try {
            setName(JSON.parse(localStorage.getItem("user")!).name);
        } catch {
            window.location.href = "/login";
        }

        getUserRole();
    }
    , [])

    return (
        <header className="flex flex-row fixed w-[90%] h-20 ml-[15%] pl-5 mr-[10%] bg-linear-to-t from-white/80 to-white/100 z-100">
            <Link
                to="/"
                className="w-70 h-10 justify-center text-center bg-palette-color4 pt-1 rounded-xl font-semibold text-xl text-white self-center select-none"
            >
                Вернуться на главную
            </Link>
            <div onClick={() => {setUserBlock(!userBlock); setUseBlockVisible(true)}} className="flex flex-row ml-auto mr-[10%]">
                <div 
                    className="w-10 h-10 select-none rotate-90 text-center self-center font-extralight text-black/50 text-2xl"
                >
                    {(userBlock ? "<" : ">")}
                </div>
                <button className="select-none rounded-4xl text-2xl font-bold bg-[#1D3266] text-white w-10 h-10 flex items-center justify-center self-center hover:cursor-pointer">
                    {
                        name[0]
                    }
                </button>
            </div>
            {/* User block drop box */}
            <div className={`w-60 bg-white drop-shadow-2xl rounded-md fixed right-8 top-17 ${userBlock ? "dropBoxAnimationIn" : "dropBoxAnimationOut"} ${userBlockVisible ? "" : "hidden"}`}>
                <nav className="flex flex-col p-5">
                    <div>{name}<br/>
                        <div className="w-[120%] h-[1px] -mx-5 my-2 bg-black"></div>
                        Роль: {userRole}<br/>
                        ID: {localStorage.getItem("userID")}
                    </div>
                    <div className="w-[120%] h-[1px] -mx-5 my-2 bg-black"></div>
                    { userRole === "boss" ? 
                        <>
                            <a href="/panel" className="select-none hover:cursor-pointer">Добавление рабочего</a>
                            <div className="w-[120%] h-[1px] -mx-5 my-2 bg-black"></div>
                        </>
                        :
                        <></>
                    }
                    <button onClick={() => {localStorage.clear(); window.location.href = "/login"}} className="w-20 h-10 bg-palette-color4 text-white hover:cursor-pointer rounded-xl">Выход</button>
                </nav>
            </div>
        </header>
    );
};

export default Header;
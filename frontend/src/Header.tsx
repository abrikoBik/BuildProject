import { useState, useEffect } from "react";
import { Link } from "react-router-dom";

const Header = () => {

    const [userBlock, setUserBlock] = useState(false);
    const [userBlockVisible, setUseBlockVisible] = useState(false);

    const [name, setName] = useState("");

    useEffect(() => {
        try {
            setName(JSON.parse(localStorage.getItem("user")!).name[0]);
        } catch {
            window.location.href = "/login";
        }
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
            <a
                href=""
                className="text-2xl self-center max-md:text-lg font-[400] ml-5 text-[#6f6d74] select-none"
            >
                О нас
            </a>
            <div onClick={() => {setUserBlock(!userBlock); setUseBlockVisible(true)}} className="flex flex-row ml-auto mr-[10%]">
                <div 
                    className="w-10 h-10 select-none rotate-90 text-center self-center font-extralight text-black/50 text-2xl"
                >
                    {(userBlock ? "<" : ">")}
                </div>
                <button className="select-none rounded-4xl text-2xl font-bold bg-[#1D3266] text-white w-10 h-10 flex items-center justify-center self-center hover:cursor-pointer">
                    {
                        name
                    }
                </button>
            </div>
            {/* User block drop box */}
            <div className={`w-60 bg-white drop-shadow-2xl rounded-md fixed right-8 top-17 ${userBlock ? "dropBoxAnimationIn" : "dropBoxAnimationOut"} ${userBlockVisible ? "" : "hidden"}`}>
                <nav className="flex flex-col p-5">
                    <a href="">Учетная запись</a>
                    <a href="">Настройки</a>
                    { localStorage.getItem("role") === "admin" ?
                    <>
                        <div className="w-[120%] h-[1px] -mx-5 my-2 bg-black"></div>
                        <a href="/panel">Редактирование рабочих</a>
                    </> :
                    <></> }
                    <div className="w-[120%] h-[1px] -mx-5 my-2 bg-black"></div>
                    <a href="">Выход</a>
                </nav>
            </div>
        </header>
    );
};

export default Header;
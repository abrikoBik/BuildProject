import { useEffect, useState } from "react";
import axios from "axios";

const ManagerPanel = () => {

    const [input, setInput] = useState({
        id: -1,
        role: "client"
    })

    const [userRole, setUserRole] = useState("");

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

    const handleWorkerAdd = async () => {
        try {
            if(input.id >= 0)
            {
                axios.post(
                    `/api/users/${input.id}/worker`,
                    {"role":input.role},
                    {
                        headers: {
                            "Content-Type": "application/json"
                        }
                    }
                );
            }

            console.log("Рабочий добавлен!");
        } catch (e : unknown) {
            console.log(e);
        }
    }

    useEffect(() => {
        getUserRole();
    }, [])

    return (
        <main>
            { userRole === "boss" ?
            <>
                <h2 className="font-bold text-2xl m-1">Панель редактирования ролей</h2>
                <div className="ml-2 mt-2">
                    <label className="mr-2 !text-black" htmlFor="id">Выберите ID пользователя</label> 
                    <input type="text" id="id" onChange={(e) => setInput({...input, id: parseInt(e.currentTarget.value)})} className="!border-1 !h-5 !w-10" /> <br/>
                    <label htmlFor="roles" className="mr-2 !text-black">Выберите роль</label>
                    <select onChange={(e) => setInput({...input, role: e.currentTarget.value})} name="roles" id="roles" className="bg-gray-300 p-1 border-2">
                        <option>boss</option>
                        <option>worker</option>
                        <option>client</option>
                    </select>
                    <br/><input onClick={handleWorkerAdd} type="button" value="Сменить роль" className="px-5 py-2 mt-2 hover:cursor-pointer hover:bg-gray-200 bg-gray-400 rounded-xl" />
                    <div>ID: {input.id < 0 ? "Не выбран!" : input.id}<br/>Role: {input.role}</div>
                </div>            
            </> :
            <div>Админские права не обнаружены!</div>
            }
        </main>
    );
}

export default ManagerPanel;
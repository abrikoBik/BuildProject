import { useState } from 'react';
import { Header } from './Header.tsx'

export const Register = () =>
{
    const [passShow, setPassShow] = useState(false);

    function passShowToggle ()
    {
        setPassShow(!passShow);
    }

    return (
        <>
            <Header/>
            
            <section className="h-[100vh] m-0 flex justify-center items-center">
                <div className="flex flex-col text-center p-3 justify-center align-center w-[400px] align-middle h-[700px] bg-gray-50 border-2 border-gray-200 rounded-sm">
                    <div className="m-5">Название приложения</div>
                    <div className="m-5 text-xl font-inter">Регистрация</div>
                    <div className="w-13 ml-10 mb-[-3%] bg-gray-50 z-10">ФИО</div>
                    <input id="name" type="text" className="text-center p-2 mx-5 mb-5 bg-gray-50 border-2 border-gray-300"/>
                    <div className="w-15 ml-10 mb-[-3%] bg-gray-50 z-10">Логин</div>
                    <input id="login" type="text" className="text-center p-2 mx-5 mb-5 bg-gray-50 border-2 border-gray-300"/>
                    <div className="w-15 ml-10 mb-[-3%] bg-gray-50 z-10">Пароль</div>
                    <input id="password" type={passShow ? "password" : "text"} className="text-center p-2 mx-5 mb-5 bg-gray-50 border-2 border-gray-300"/>
                    <button onClick={() => {passShowToggle()}} className="text-2xl ml-70 mt-[-14.5%] hover:cursor-pointer"><link rel="stylesheet" href={passShow ? "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@24,400,0,0&icon_names=visibility" : "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@24,400,0,0&icon_names=visibility_off"} /><span className="material-symbols-outlined">
{passShow ? "visibility" : "visibility_off"}
</span></button>
                    <div className="mt-5 w-38 ml-10 mb-[-3%] bg-gray-50 z-10">Повторить пароль</div>
                    <input id="password" type={passShow ? "password" : "text"} className="text-center p-2 mx-5 mb-5 bg-gray-50 border-2 border-gray-300"/>
                    <button onClick={() => {passShowToggle()}} className="text-2xl ml-70 mt-[-14.5%] hover:cursor-pointer"><link rel="stylesheet" href={passShow ? "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@24,400,0,0&icon_names=visibility" : "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@24,400,0,0&icon_names=visibility_off"} /><span className="material-symbols-outlined">
{passShow ? "visibility" : "visibility_off"}
</span></button>
                    <select className="mt-5 mb-5 mx-25 p-2 text-center text-lg rounded-sm bg-gray-50 border-2 border-gray-200">
                        <option>Подрядчик</option>
                        <option>Заказчик</option>
                    </select>
                    <a href="/login" className="my-1 mx-25 p-1 px-3 rounded-sm bg-blue-200 hover:cursor-pointer">Регистрация</a>
                    <div className="m-5 flex flex-row justify-center">
                        <div className="mx-5 my-3 bg-gray-200 w-full h-[2px]"></div>
                        <div>или</div>
                        <div className="mx-5 my-3 bg-gray-200 w-full h-[2px]"></div>
                    </div>
                    <a href="/register" className="my-1 mx-25 p-1 px-3 rounded-sm bg-gray-200 hover:cursor-pointer">Вход</a>
                </div>
            </section>
        </>
    );
};
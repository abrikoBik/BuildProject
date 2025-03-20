import { useState } from "react";
import { Header } from "./Header.tsx";

export const Login = () => {
    const [passShow, setPassShow] = useState(false);
    const passShowToggle = () => setPassShow(!passShow);

    return (
        <>
            <Header />
            <section className="h-[100vh] m-0 flex justify-center items-center drop-shadow-md">
                <div className="flex flex-col text-center p-3 justify-center align-center w-[400px] align-middle h-[500px] bg-palette-color2 border-2 border-palette-color5 rounded-sm">
                    <div className="text-3xl font-bold m-5">StruKtura</div>
                    <div className="w-full m-0 h-[2px] bg-palette-color3"></div>
                    <div className="m-5 text-2xl">Вход</div>
                    <div className="w-15 ml-10 mb-[-3%] bg-palette-color2 z-10">
                        Логин
                    </div>
                    <input
                        id="login"
                        type="text"
                        className="text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-palette-color5 rounded-sm"
                    />
                    <div className="w-18 ml-10 mb-[-3%] bg-palette-color2 z-10">
                        Пароль
                    </div>
                    <input
                        id="password"
                        type={passShow ? "password" : "text"}
                        className="text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-palette-color5 rounded-sm"
                    />
                    <button
                        onClick={() => {
                            passShowToggle();
                        }}
                        className="text-2xl ml-70 mt-[-14.5%] hover:cursor-pointer"
                    >
                        <link
                            rel="stylesheet"
                            href={
                                passShow
                                    ? "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@24,400,0,0&icon_names=visibility"
                                    : "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@24,400,0,0&icon_names=visibility_off"
                            }
                        />
                        <span className="material-symbols-outlined">
                            {passShow ? "visibility" : "visibility_off"}
                        </span>
                    </button>
                    <a
                        href="/login"
                        className="mt-9 mb-1 mx-25 p-1 px-3 rounded-sm bg-palette-color3 font-bold hover:cursor-pointer"
                    >
                        Вход
                    </a>
                    <div className="m-5 flex flex-row justify-center">
                        <div className="mx-5 my-3 bg-palette-color3 w-full h-[2px]"></div>
                        <div>или</div>
                        <div className="mx-5 my-3 bg-palette-color3 w-full h-[2px]"></div>
                    </div>
                    <a
                        href="/register"
                        className="my-1 mx-25 p-1 px-3 rounded-sm bg-palette-color4 font-bold hover:cursor-pointer"
                    >
                        Регистрация
                    </a>
                </div>
            </section>
        </>
    );
};

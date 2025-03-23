import React, { useState } from "react";
import axios, { AxiosResponse } from "axios";
import { Header } from "./Header.tsx";
import { Link } from 'react-router-dom'
import HalfImage from './assets/auth_image.png'

interface User {
    login: string;
    name: string;
    email: string;
    password: string;
}

export const Register = () => {
//    const [passShow, setPassShow] = useState(false);
//    const passShowToggle = () => setPassShow(!passShow);

    const [post, setPost] = useState<User>({
        login: "",
        name: "",
        email: "",
        password: "",
    });

    const handleInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        if(event.target.name != "checkPass")
        {
            setPost({ ...post, [event.target.name]: event.target.value });
        }
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        try {
            const data: AxiosResponse = await axios.post<User>(
                "http://0.0.0.0:3000/users",
                post,
            );
            console.log(data);
        } catch (err) {
            console.error(err);
        }
    };

    return (
        <section className="flex flex-row">
            <section className="flex flex-col w-[50vw] h-[100vh]">
                {/* Header */}
                <Header/>

                {/* Register inputs */}
                <section className="w-[50vw] h-[100vh]">
                        <div className="flex items-center justify-center h-[100vh]">
                            <form onSubmit={handleSubmit} className="flex flex-col">
                                <label htmlFor="name" >ФИО</label>
                                <input 
                                    onChange={handleInput}
                                    type="text"
                                    name="name"
                                />
                                <label htmlFor="login" >Логин</label>
                                <input
                                    onChange={handleInput} 
                                    type="text"
                                    name="login"
                                />
                                <label htmlFor="login" >Email</label>
                                <input
                                    onChange={handleInput} 
                                    type="text"
                                    name="email"
                                />
                                <label htmlFor="password" >Пароль</label>
                                <input 
                                    onChange={handleInput}
                                    type="text"
                                    name="password"
                                />
                                <label htmlFor="checkPass">Повторите пароль</label>
                                <input 
                                    onChange={handleInput}
                                    type="text"
                                    name="checkPass"
                                />
                                <div className="flex flex-row justify-start w-full mb-5 mt-[-2vh]">
                                    <input type="checkbox" name="savePass" className="m-0 w-5 h-5 border-2 border-[#2B4074]"/>
                                    <label htmlFor="savePass" className="ml-1 text-md">Даю согласие на обработку персональных данных</label>
                                </div>
                                <button className="mb-5 bg-[#2B4074] text-white h-[40px] w-[40%] self-center font-[600] rounded-[3px]">Регистрация</button>
                                <Link to="/login" className="mb-5 pt-1.5 border-[2px] border-[#2B4074] h-[40px] w-[40%] self-center font-[600] rounded-[3px] text-center">Вход</Link>
                            </form>
                        </div>
                </section>
            </section>

            <img src={HalfImage} alt="half_image" className="w-[50vw] h-[100vh] max-2xl:hidden" />
            {/*<section className="flex flex-row">
                <form
                    onSubmit={handleSubmit}
                    className="flex flex-col text-center p-3 justify-center align-center w-[400px] align-middle h-[700px] bg-palette-color2 border-2 border-palette-color5 rounded-sm"
                >
                    <div className="text-3xl font-bold m-5">StruKtura</div>
                    <div className="w-full m-0 h-[2px] bg-palette-color3"></div>
                    <div className="m-5 text-2xl">Регистрация</div>

                    <div className="w-13 ml-10 mb-[-4%] bg-palette-color2 z-10">
                        ФИО
                    </div>
                    <input
                        onChange={handleInput}
                        name="name"
                        type="text"
                        className="select-none text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-paletter-color5"
                    />

                    <div className="w-15 ml-10 mb-[-3%] bg-palette-color2 z-10">
                        Логин
                    </div>
                    <input
                        onChange={handleInput}
                        name="login"
                        type="text"
                        className="select-none text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-palette-color5"
                    />

                    <div className="w-15 ml-10 mb-[-3%] bg-palette-color2 z-10">
                        Email
                    </div>
                    <input
                        onChange={handleInput}
                        name="email"
                        type="text"
                        className="select-none text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-palette-color5"
                    />

                    <div className="w-18 ml-10 mb-[-3%] bg-palette-color2 z-10">
                        Пароль
                    </div>
                    <input
                        onChange={handleInput}
                        name="password"
                        type={passShow ? "text" : "password"}
                        className="select-none text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-palette-color5"
                    />

                    <div className="mt-5 w-38 ml-10 mb-[-3%] bg-palette-color2 z-10">
                        Повторить пароль
                    </div>
                    <input
                        //onChange={handleInput}
                        name="confirmPassword"
                        type={passShow ? "text" : "password"}
                        className="select-none text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-palette-color5"
                    />

                    <button
                        type="button"
                        onClick={passShowToggle}
                        className="text-2xl hover:cursor-pointer"
                    >
                        <span className="material-symbols-outlined">
                            {passShow ? "visibility_off" : "visibility"}
                        </span>
                    </button>

                    <div className="m-5 flex flex-row justify-center">
                        <div className="mx-5">
                            <input
                                name="role"
                                type="radio"
                                value="contractor"
                                //checked={post.role === "contractor"}
                                //onChange={handleInput}
                            />
                            <label className="m-2">Подрядчик</label>
                        </div>
                        <div className="mx-5">
                            <input
                                name="role"
                                type="radio"
                                value="client"
                                //checked={post.role === "client"}
                                //onChange={handleInput}
                            />
                            <label className="m-2">Заказчик</label>
                        </div>
                    </div>

                    <button
                        type="submit"
                        className="my-1 mx-25 p-1 px-3 font-bold rounded-sm bg-palette-color4 hover:cursor-pointer"
                    >
                        Регистрация
                    </button>

                    <div className="m-5 flex flex-row justify-center">
                        <div className="mx-5 my-3 bg-palette-color3 w-full h-[2px]"></div>
                        <div>или</div>
                        <div className="mx-5 my-3 bg-palette-color3 w-full h-[2px]"></div>
                    </div>
                    <a
                        href="/login"
                        className="my-1 mx-25 p-1 px-3 font-bold rounded-sm bg-palette-color3 hover:cursor-pointer"
                    >
                        Вход
                    </a>
                </form>
            </section>*/}
        </section>
    );
};

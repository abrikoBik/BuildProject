import React, { useState } from "react";
import axios, { AxiosError, AxiosResponse } from "axios";
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

            alert(`Пользователь ${post.name} был успешно зарегестрирован!`);

            console.log(data);
        } catch (err : unknown) {
            if(err instanceof AxiosError)
            {
                alert(err?.response?.data);
            }
            console.error(err);
        }
    };

    return (
        <section className="flex flex-row overflow-hidden">
            <section className="flex flex-col w-[50vw] h-[100vh]">
                {/* Header */}
                <Header/>

                {/* Register inputs */}
                <section className="lg:w-[50vw] w-[100vw] h-[95vh]">
                        <div className="flex items-center justify-center h-[100vh] ">
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
                                <div className="flex flex-row justify-start w-full mb-5 mt-[-1vh]">
                                    <input type="checkbox" name="savePass" className="m-0 w-5 h-5 border-2 border-[#2B4074]"/>
                                    <label htmlFor="savePass" className="ml-1 text-md">Даю согласие на обработку персональных данных</label>
                                </div>
                                <button className="mb-5 bg-[#2B4074] text-white max-lg:h-[30px] h-[40px] w-[30%] self-center font-[600] rounded-[3px]">Регистрация</button>
                                <Link to="/login" className="mb-5 border-[2px] border-[#2B4074] max-lg:h-[26px] h-[36px] w-[30%] self-center font-[600] rounded-[3px] text-center">Вход</Link>
                            </form>
                        </div>
                </section>
            </section>

            <img src={HalfImage} alt="half_image" className="w-[100%] h-[100vh] object-cover max-lg:hidden" />
        </section>
    );
};

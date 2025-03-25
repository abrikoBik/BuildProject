import React, { useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import HalfImage from "./assets/auth_image.png";
import errors from "./util/errors.ts";
import { UserRequest } from "./types/user.ts";
import Logo from './assets/logo.svg'

export const Register = () => {
    //    const [passShow, setPassShow] = useState(false);
    //    const passShowToggle = () => setPassShow(!passShow);

    const [userRequest, setUserRequest] = useState<UserRequest>({
        login: "",
        name: "",
        email: "",
        password: "",
    });

    const handleInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.name != "checkPass") {
            setUserRequest({
                ...userRequest,
                [event.target.name]: event.target.value,
            });
        }
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        try {
            const { data, status } = await axios.post<UserRequest>(
                "http://0.0.0.0:3000/users",
                userRequest,
            );
            if (status != 200) {
                alert(data);
                return;
            }

            alert(`Регистрация прошла успешно!`);
            console.log(data);
        } catch (err: unknown) {
            errors.axiosError(err);
        }
    };

    return (
        <section className="flex flex-row overflow-hidden">
            <section className="flex flex-col w-[50vw] h-[100vh]">
                {/* Header */}
                <header className="flex flex-row w-[100vw] lg:w-[50vw] h-[5vh] pt-2 px-5">
                    <img src={Logo} alt="logo" />
                    <section className="flex flex-row w-full mt-1 ml-[-73px] justify-center">
                        <a href="" className="text-xl max-md:text-lg font-[400] mr-5 text-[#6f6d74]">
                            Главная
                        </a>
                        <a href="" className="text-xl max-md:text-lg font-[400] mr-5 text-[#6f6d74]">
                            О нас
                        </a>
                    </section>
                </header>

                {/* Register inputs */}
                <section className="max-md:px-[5vw] lg:w-[50vw] w-[100vw] h-[95vh]">
                    <div className="flex items-center justify-center h-[100vh] ">
                        <form onSubmit={handleSubmit} className="flex flex-col">
                            <label htmlFor="name">ФИО</label>
                            <input
                                onChange={handleInput}
                                type="text"
                                name="name"
                            />
                            <label htmlFor="login">Логин</label>
                            <input
                                onChange={handleInput}
                                type="text"
                                name="login"
                            />
                            <label htmlFor="login">Email</label>
                            <input
                                onChange={handleInput}
                                type="text"
                                name="email"
                            />
                            <label htmlFor="password">Пароль</label>
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
                                <input
                                    type="checkbox"
                                    name="savePass"
                                    className="m-0 w-5 h-5 border-2 border-[#2B4074]"
                                />
                                <label
                                    htmlFor="savePass"
                                    className="ml-1 text-md"
                                >
                                    Даю согласие на обработку персональных
                                    данных
                                </label>
                            </div>
                            <button className="mb-5 bg-[#2B4074] text-white max-lg:h-[30px] h-[40px] w-[30%] self-center font-[600] rounded-[3px]">
                                Регистрация
                            </button>
                            <Link
                                to="/login"
                                className="mb-5 lg:pt-1 border-[2px] border-[#2B4074] max-lg:h-[26px] h-[36px] w-[30%] self-center font-[600] rounded-[3px] text-center"
                            >
                                Вход
                            </Link>
                        </form>
                    </div>
                </section>
            </section>

            <img
                src={HalfImage}
                alt="half_image"
                className="w-[100%] h-[100vh] object-cover max-lg:hidden"
            />
        </section>
    );
};

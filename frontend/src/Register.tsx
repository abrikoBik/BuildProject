import React, { useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import HalfImage from "./assets/auth_image.png";
import errors from "./util/errors.ts";
import { UserRequest } from "./types/user.ts";
import Logo from "./assets/logo.svg";

const Register = () => {
    const [visibility, setVisibility] = useState(true);
    const toggleVisibility = () => {
        setVisibility(!visibility);
    };
    const [visibilityC, setVisibilityC] = useState(true);
    const toggleVisibilityC = () => {
        setVisibilityC(!visibilityC);
    };
    const [checkPass, setCheckPass] = useState("");

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
        } else {
	    setCheckPass(event.target.value);
	}
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        try {
	    if(checkPass != userRequest.password)
	    {
		alert("Пароли не совпадают!");
		return;
            }
            const { data, status } = await axios.post<UserRequest>(
                "http://0.0.0.0:3000/users",
                userRequest,
            );
            if (status != 200) {
                alert(data);
                return;
            }

            alert("Регистрация прошла успешно!");
            console.log(data);
        } catch (err: unknown) {
            errors.axiosError(err);
        }
    };

    return (
        <section className="flex h-screen">
            <div className="flex-1 flex flex-col items-center justify-center p-8 max-w-[50vw] max-lg:max-w-full">
                <header className="flex w-1/2 max-lg:w-full absolute top-0 left-0 p-4">
                    <img src={Logo} alt="logo" className="h-8" />
                    <nav className="flex items-center justify-center w-full gap-5 ml-[-4vw]">
                        <Link
                            to="/"
                            className="text-lg text-palette-color5 hover:cursor-pointer transition-colors"
                        >
                            Главная
                        </Link>
                        <a className="text-lg text-palette-color5 hover:cursor-pointer transition-colors">
                            О нас
                        </a>
                    </nav>
                </header>

                <form
                    onSubmit={handleSubmit}
                    className="w-full max-w-[400px] flex flex-col gap-6"
                >
                    <div className="flex flex-col gap-0">
                        <label className="flex items-center gap-2">
                            <span className="material-symbols-outlined">
                                person
                            </span>
                            Логин
                        </label>
                        <input
                            onChange={handleInput}
                            type="text"
                            name="login"
                            className="w-full"
                        />
                    </div>

                    <div className="flex flex-col gap-0">
                        <label className="flex items-center gap-2">
                            <span className="material-symbols-outlined">
                                badge
                            </span>
                            ФИО
                        </label>
                        <input
                            onChange={handleInput}
                            type="text"
                            name="name"
                            className="w-full"
                        />
                    </div>

                    <div className="flex flex-col gap-0">
                        <label className="flex items-center gap-2">
                            <span className="material-symbols-outlined">
                                mail
                            </span>
                            Email
                        </label>
                        <input
                            onChange={handleInput}
                            type="text"
                            name="email"
                            className="w-full"
                        />
                    </div>

                    <div className="flex flex-col gap-0 relative">
                        <label className="flex items-center gap-2">
                            <span className="material-symbols-outlined">
                                key
                            </span>
                            Пароль
                        </label>
                        <input
                            onChange={handleInput}
                            type={visibility ? "password" : "text"}
                            name="password"
                            className="w-[90%]"
                        />
                        <div className="bg-palette-color4 absolute bottom-0 right-0 w-[10%] h-[1px]"></div>
                        <button
                            type="button"
                            onClick={toggleVisibility}
                            className="absolute right-2 bottom-1 text-gray-500 hover:text-gray-700"
                        >
                            <span className="material-symbols-outlined text-xl">
                                {visibility ? "visibility" : "visibility_off"}
                            </span>
                        </button>
                    </div>

                    <div className="flex flex-col gap-0 relative">
                        <label className="flex items-center gap-2">
                            <span className="material-symbols-outlined">
                                key
                            </span>
                            Повторите пароль
                        </label>
                        <input
                            onChange={handleInput}
                            type={visibilityC ? "password" : "text"}
                            name="checkPass"
                            className="w-[90%]"
                        />
                        <div className="bg-palette-color4 absolute bottom-0 right-0 w-[10%] h-[1px]"></div>
                        <button
                            type="button"
                            onClick={toggleVisibilityC}
                            className="absolute right-2 bottom-1 text-gray-500 hover:text-gray-700"
                        >
                            <span className="material-symbols-outlined text-xl">
                                {visibilityC ? "visibility" : "visibility_off"}
                            </span>
                        </button>
                    </div>

                    <div className="flex justify-between items-center">
                        <label className="flex items-center gap-2 text-xs my-2">
                            <input
                                type="checkbox"
                                name="savePass"
                                className="w-5 h-5 accent-palette-color4"
                            />
                            Даю согласие на обработку персональных данных
                        </label>
                    </div>

                    <button
                        type="submit"
                        className="w-1/2 mb-[-1vh] bg-palette-color4 border-2 border-palette-color5 text-palette-color1 py-2 rounded-t-[5px] font-semibold cursor-pointer self-center self-justify-center"
                    >
                        Регистрация
                    </button>

                    <Link
                        to="/login"
                        className="w-1/2 text-center border-2 border-palette-color5 py-2 rounded-b-[5px] font-semibold text-palette-color4 self-center self-justify-center"
                    >
                        Вход
                    </Link>
                </form>
            </div>

            <div className="flex-1 max-lg:hidden">
                <img
                    src={HalfImage}
                    alt="half_image"
                    className="w-full h-screen object-cover"
                />
            </div>
        </section>
    );
};

export default Register;

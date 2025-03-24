import HalfImage from "./assets/auth_image.png";
import { useState } from "react";
import { useCookies } from "react-cookie";
import { Link } from "react-router-dom";
import { Header } from "./Header.tsx";
import axios, { AxiosResponse } from "axios";
import errors from "./util/errors.ts";
import { SessionRequest, SessionResponse } from "./types/session.ts";
import { UserResponse } from "./types/user.ts";

export const Login = () => {
    const [sessionRequest, setSessionRequest] = useState<SessionRequest>({
        login: "",
        password: "",
    });

    const setCookie = useCookies(["access-token"])[1];

    const handleInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        setSessionRequest({
            ...sessionRequest,
            [event.target.name]: event.target.value,
        });
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        try {
            const session: AxiosResponse<SessionResponse> = await axios.post(
                "http://0.0.0.0:3000/sessions",
                sessionRequest,
                { withCredentials: true },
            );
            setCookie("access-token", session.data.token, {
                path: "/",
                sameSite: "none",
                secure: true,
            });

            const user: AxiosResponse<UserResponse> = await axios.get(
                `http://0.0.0.0:3000/users/${session.data.userID}`,
            );
            if (user.status != 200) {
                alert(user.data);
                return;
            }
            localStorage.setItem("user", JSON.stringify(user.data));

            const storedUser: UserResponse = JSON.parse(
                localStorage.getItem("user")!,
            );
            alert(`Добро пожаловать, ${storedUser.name}!`);
        } catch (err: unknown) {
            errors.axiosError(err);
        }
    };

    return (
        <section className="flex flex-row">
            <section className="flex flex-col w-[50vw] h-[100vh]">
                {/* Header */}
                <Header />

                {/* Login inputs */}
                <section className="lg:w-[50vw] w-[100vw] h-[95vh]">
                    <div className="flex items-center justify-center h-[100vh]">
                        <form onSubmit={handleSubmit} className="flex flex-col">
                            <label htmlFor="email" className="">
                                Логин
                            </label>
                            <input
                                onChange={handleInput}
                                type="text"
                                name="login"
                            />
                            <label htmlFor="password">Пароль</label>
                            <input
                                onChange={handleInput}
                                type="text"
                                name="password"
                            />
                            <div className="flex flex-row justify-start w-full mb-5 mt-[-2vh]">
                                <input
                                    type="checkbox"
                                    name="savePass"
                                    className="m-0 w-5 h-5 border-2 border-[#2B4074]"
                                />
                                <label htmlFor="savePass" className="ml-1">
                                    Сохранить вход
                                </label>
                                <a
                                    href=""
                                    className="ml-[50%] self-end text-[#2B4074]"
                                >
                                    Забыли пароль?
                                </a>
                            </div>
                            <button className="mb-5 bg-[#2B4074] text-white max-lg:h-[30px] h-[40px] w-[40%] self-center font-[600] rounded-[3px]">
                                Вход
                            </button>
                            <Link
                                to="/register"
                                className="mb-5 pt-1.5 border-[2px] border-[#2B4074] max-lg:h-[30px] h-[40px] w-[40%] self-center font-[600] rounded-[3px] text-center"
                            >
                                Регистрация
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

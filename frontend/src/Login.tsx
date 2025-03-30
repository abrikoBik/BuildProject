import HalfImage from "./assets/auth_image.png";
import { useState } from "react";
import { useCookies } from "react-cookie";
import { Link } from "react-router-dom";
import axios, { AxiosResponse } from "axios";
import errors from "./util/errors.ts";
import { SessionRequest, SessionResponse } from "./types/session.ts";
import { UserResponse } from "./types/user.ts";
import Logo from "./assets/logo.svg";

export const Login = () => {
    const [visibility, setVisibility] = useState(true);

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
    
    const toggleVisibility = () => {
        setVisibility(!visibility);
    };

    return (
        <section className="flex flex-row overflow-hidden">
            <section className="flex flex-col max-lg:w-[100vw] w-[50vw] h-[100vh]">
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
                                className="mb-[4vh]"
                            />
                            <div>
                                <label htmlFor="password">Пароль</label>
                                <input
                                    onChange={handleInput}
                                    type={visibility ? "password" : "text"}
                                    name="password"
                                    className="mb-[2vh]"
                                />
                                <span onClick={toggleVisibility} className="material-symbols-outlined absolute translate-x-[-100%] translate-y-1/3 select-none cursor-pointer">{visibility ? "visibility" : "visibility_off"}</span>
                            </div>
                            <div className="flex flex-row justify-start w-full mb-5">
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
                                    className="ml-[5vw] self-end text-[#2B4074]"
                                >
                                    Забыли пароль?
                                </a>
                            </div>
                            <button className="mb-5 bg-[#2B4074] text-white max-lg:h-[30px] h-[40px] w-[40%] self-center font-[600] rounded-[3px]">
                                Вход
                            </button>
                            <Link
                                to="/register"
                                className="mb-5 lg:pt-1.5 border-[2px] border-[#2B4074] max-lg:h-[30px] h-[40px] w-[40%] self-center font-[600] rounded-[3px] text-center"
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

import HalfImage from './assets/auth_image.png'
import Logo from './assets/logo.svg'
//import { Header } from "./Header.tsx";

export const Login = () => {
    return (
        <section className="flex flex-row">
            <section className="flex flex-col w-[50vw] h-[100vh]">
                {/* Header */}
                <section className="flex flex-row w-[50vw] h-[5vh] pt-2">
                    <img src={Logo} alt="logo"/>
                    <section className="flex flex-row ml-65 mt-3 self-center">
                        <a href="" className="text-xl font-[400] mr-5 text-[#6f6d74]">Главная</a>
                        <a href="" className="text-xl font-[400] mr-5 text-[#6f6d74]">О нас</a>
                    </section>
                </section>

                {/* Login inputs */}
                <div className="w-[50vw] h-[100vh]">
                    <div className="flex items-center justify-center pt-[35vh]">
                        <form action="submit" className="flex flex-col">
                            <label htmlFor="email" className="">Email</label>
                            <input type="text" name="email"/>
                            <label htmlFor="password">Пароль</label>
                            <input type="text" name="password"/>
                            <div className="flex flex-row justify-start w-full mb-5 mt-[-2vh]">
                                <input type="checkbox" name="savePass" className="m-0 w-5 h-5 border-2 border-[#2B4074]"/>
                                <label htmlFor="savePass" className="ml-1">Сохранить вход</label>
                                <a href="" className="ml-[50%] self-end text-[#2B4074]">Забыли пароль?</a>
                            </div>
                            <button className="mb-5 bg-[#2B4074] text-white h-[40px] w-[40%] self-center font-[600] rounded-[3px]">Вход</button>
                            <button className="mb-5 border-[2px] border-[#2B4074] h-[40px] w-[40%] self-center font-[600] rounded-[3px]">Регистрация</button>
                        </form>
                    </div>
                </div>
            </section>

            <img src={HalfImage} alt="half_image" className="w-[50vw] h-[100vh]" />
        </section>
    );
};

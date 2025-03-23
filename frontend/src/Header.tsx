//import { Link } from "react-router-dom";
import Logo from "./assets/logo.svg";

export const Header = () => {
    return (
        <header className="flex flex-row w-[50vw] h-[5vh] pt-2 pl-5">
                <img src={Logo} alt="logo"/>
                <section className="flex flex-row w-full mt-3 justify-center">
                    <a href="" className="text-xl font-[400] mr-5 text-[#6f6d74]">Главная</a>
                    <a href="" className="text-xl font-[400] mr-5 text-[#6f6d74]">О нас</a>
                </section>
        </header>
    );
};

import Logo from "./assets/logo_2.svg";
// import { useState } from "react";

const Sidebar = ({ show, choice }: { show: boolean, choice: number }) => {
    return (
        <section className={`fixed top-0 left-0 flex flex-col items-center h-screen w-[15vw] mr-6 bg-palette-color4 max-md:hidden max-md:m-0 ${show ? "" : "hidden"}`}>
            <img src={Logo} alt="logo" className="mt-6 mx-2 mb-4 h-14" />
            <nav className="flex flex-col w-full m-0">
                <a
                    href="/projects"
                    className={`w-full text-lg font-semibold py-6 border-b-1 border-t-1 border-palette-color1 text-center text-palette-color1 hover:bg-[#232c33] transition-colors ${choice == 1 ? "bg-[#5b7385]" : "bg-palette-color4"}`}
                >
                    Проекты
                </a>
                <a
                    href="/estimates"
                    className={`w-full text-lg font-semibold py-6 mt-[-1px] border-b-1 border-t-1 border-palette-color1 text-center text-palette-color1 hover:bg-[#232c33] transition-colors ${choice == 2 ? "bg-[#5b7385]" : "bg-palette-color4"}`}
                >
                    Сметы
                </a>
                <a
                    href="/blueprints"
                    className={`w-full text-lg font-semibold py-6 mt-[-1px] border-b-1 border-t-1 border-palette-color1 text-center text-palette-color1 hover:bg-[#232c33] transition-colors ${choice == 3 ? "bg-[#5b7385]" : "bg-palette-color4"}`}
                >
                    Чертежи
                </a>
            </nav>
        </section>
    );
};

export default Sidebar;

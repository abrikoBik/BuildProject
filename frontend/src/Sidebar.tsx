import Logo from "./assets/logo_2.svg"

export const Sidebar = () => {
    return (
        <section className="flex flex-col items-center h-screen w-[15vw] mr-6 bg-palette-color4">
            <img src={Logo} alt="logo" className="mt-8 mx-2 mb-4 h-14" />
            <nav className="flex flex-col w-full m-0">
                <a href="" className="w-full py-4 border-b-1 border-t-1 border-palette-color1 text-center text-palette-color1 hover:bg-[#232c33] transition-colors">Проекты</a>
                <a href="" className="w-full py-4 mt-[-1px] border-b-1 border-t-1 border-palette-color1 text-center text-palette-color1 hover:bg-[#232c33] transition-colors">Сметы</a>
                <a href="" className="w-full py-4 mt-[-1px] border-b-1 border-t-1 border-palette-color1 text-center text-palette-color1 hover:bg-[#232c33] transition-colors">Чертежи</a>
            </nav>
        </section>
    );
}
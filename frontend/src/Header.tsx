import { Link } from "react-router-dom";

const Header = () => {
    return (
        <header className="flex flex-row w-[80%] h-20 ml-[17%] mr-[10%]">
            <Link
                to="/"
                className="w-80 h-10 justify-center text-center border-2 bg-palette-color4 pt-1 rounded-xl font-semibold text-xl text-white self-center"
            >
                Вернуться на главную
            </Link>
            <a
                href=""
                className="text-2xl self-center max-md:text-lg font-[400] ml-5 text-[#6f6d74]"
            >
                О нас
            </a>
            <div className="ml-auto pointer-events-none select-none rounded-4xl text-2xl font-bold bg-[#1D3266] text-white w-10 h-10 flex items-center justify-center self-center">
                V
            </div>
        </header>
    );
};

export default Header;
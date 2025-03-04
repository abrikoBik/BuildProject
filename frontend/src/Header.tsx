import { Link } from 'react-router-dom'
import ReactLogo from './assets/react.svg'

export const Header = () =>
{
    return (
        <header className="fixed top-0 left-0 right-0 z-50 bg-gray-100 bg-opacity-75 backdrop-blur-sm">
            <div className="mx-auto flex items-center justify-between p-2">
                <Link to="/" className="sm:mr-10 md:mr-40 flex-shrink-0"><img src={ReactLogo} alt="logo" /></Link>
                
                <section className="flex flex-row space-x-4 flex-grow justify-center">
                    <div className="hover:cursor-pointer">Главная</div>
                    <div className="hover:cursor-pointer">Отзывы</div>
                    <div className="hover:cursor-pointer">О нас</div>
                </section>

                <section className="flex flex-row space-x-4 flex-shrink-0">
                    <Link to="/login" className="p-1 px-3 rounded-sm bg-gray-200 hover:cursor-pointer">Вход</Link>
                    <Link to="/register" className="p-1 px-3 rounded-sm bg-[#C598E0] hover:cursor-pointer">Регистрация</Link>
                </section>
            </div>
        </header>
    );
};

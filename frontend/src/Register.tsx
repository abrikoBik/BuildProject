import { useState } from 'react';
import { Header } from './Header.tsx'
import axios from 'axios'

export const Register = () => {
    const [passShow, setPassShow] = useState(false);
    const [post, setPost] = useState({
        login: "",
        name: "",
        email: "",
        password: "",
    })

    function passShowToggle() {
        setPassShow(!passShow);
    }

    // TODO: specify correct type
    const handleInput = (event: any) => {
        setPost({...post, [event.target.name]: event.target.value});
    };

    function handleSubmit(event: any) {
        event.preventDefault();
        
        axios.post('http://0.0.0.0:3000/register', post)
            .then(response => console.log(response))
            .catch(err => console.log(err));
    }

    return (
        <>
            <Header/>
            
            <section className="h-[100vh] m-0 flex justify-center items-center">
                <form onSubmit={handleSubmit} className="flex flex-col text-center p-3 justify-center align-center w-[400px] align-middle h-[700px] bg-palette-color2 border-2 border-palette-color5 rounded-sm">
                    <div className="text-3xl font-bold m-5">StruKtura</div>
                    <div className="w-full m-0 h-[2px] bg-palette-color3"></div>
                    <div className="m-5 text-2xl">Регистрация</div>
                    
                    {/* ФИО */}
                    <div className="w-13 ml-10 mb-[-4%] bg-palette-color2 z-10">ФИО</div>
                    <input 
                        onChange={handleInput}
                        name="name"
                        type="text" 
                        className="select-none text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-paletter-color5"
                    />
                    
                    {/* Логин */}
                    <div className="w-15 ml-10 mb-[-3%] bg-palette-color2 z-10">Логин</div>
                    <input 
                        onChange={handleInput}
                        name="login" 
                        type="text" 
                        className="select-none text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-palette-color5"
                    />

                    {/* Email */}
                    <div className="w-15 ml-10 mb-[-3%] bg-palette-color2 z-10">Email</div>
                    <input 
                        onChange={handleInput}
                        name="email" 
                        type="text" 
                        className="select-none text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-palette-color5"
                    />
                    
                    {/* Пароль */}
                    <div className="w-18 ml-10 mb-[-3%] bg-palette-color2 z-10">Пароль</div>
                    <input 
                        onChange={handleInput}
                        name="password"
                        type={passShow ? "text" : "password"} 
                        className="select-none text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-palette-color5"
                    />
                    
                    {/* Подтверждение пароля */}
                    <div className="mt-5 w-38 ml-10 mb-[-3%] bg-palette-color2 z-10">Повторить пароль</div>
                    <input 
                        //onChange={handleInput}
                        name="confirmPassword"
                        type={passShow ? "text" : "password"} 
                        className="select-none text-center p-2 mx-5 mb-5 bg-palette-color2 border-2 border-palette-color5"
                    />
                    
                    {/* Кнопка показа пароля */}
                    <button 
                        type="button"
                        onClick={passShowToggle} 
                        className="text-2xl hover:cursor-pointer"
                    >
                        <span className="material-symbols-outlined">
                            {passShow ? "visibility_off" : "visibility"}
                        </span>
                    </button>
                    {/* Радио-кнопки */}
                    <div className="m-5 flex flex-row justify-center">
                        <div className="mx-5">
                            <input 
                                name="role"
                                type="radio" 
                                value="contractor"
                                //checked={post.role === "contractor"}
                                //onChange={handleInput}
                            />
                            <label className="m-2">Подрядчик</label>
                        </div>
                        <div className="mx-5">
                            <input 
                                name="role"
                                type="radio" 
                                value="client"
                                //checked={post.role === "client"}
                                //onChange={handleInput}
                            />
                            <label className="m-2">Заказчик</label>
                        </div>
                    </div>

                    <button type="submit" className="my-1 mx-25 p-1 px-3 font-bold rounded-sm bg-palette-color4 hover:cursor-pointer">
                        Регистрация
                    </button>
                    
                    <div className="m-5 flex flex-row justify-center">
                        <div className="mx-5 my-3 bg-palette-color3 w-full h-[2px]"></div>
                        <div>или</div>
                        <div className="mx-5 my-3 bg-palette-color3 w-full h-[2px]"></div>
                    </div>
                    <a href="/login" className="my-1 mx-25 p-1 px-3 font-bold rounded-sm bg-palette-color3 hover:cursor-pointer">
                        Вход
                    </a>
                </form>
            </section>
        </>
    );
};

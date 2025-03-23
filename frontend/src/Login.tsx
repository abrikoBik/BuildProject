import HalfImage from './assets/auth_image.png'
import { useState } from 'react'
import { Link } from 'react-router-dom'
import { Header } from "./Header.tsx";
import axios, {AxiosResponse} from 'axios'

interface User {
    login: string;
    password: string;
}

export const Login = () => {

    const [post, setPost] = useState<User>({
        login: "",
        password: "",
    });
    
    const handleInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPost({ ...post, [event.target.name]: event.target.value });
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        try {
            const data: AxiosResponse = await axios.post<User>(
                "http://0.0.0.0:3000/sessions",
                post,
                {withCredentials: true}
            );
            
            axios.get(`http://0.0.0.0:3000/users/${data.data.userID}`)
            .then((resp) =>
            {
                if(resp.status == 200)
                {
                    alert(`Welcome ${resp.data.name}`);
                }
            });

            console.log(data);
        } catch (err) {
            console.error(err);
        }
    };

    return (
        <section className="flex flex-row">
            <section className="flex flex-col w-[50vw] h-[100vh]">
                {/* Header */}
                <Header/>

                {/* Login inputs */}
                <section className="w-[50vw] h-[100vh]">
                    <div className="flex items-center justify-center h-[100vh]">
                        <form onSubmit={handleSubmit} className="flex flex-col">
                            <label htmlFor="email" className="">Логин</label>
                            <input 
                                onChange={handleInput}
                                type="text"
                                name="login"/>
                            <label htmlFor="password">Пароль</label>
                            <input 
                                onChange={handleInput}
                                type="text"
                                name="password"/>
                            <div className="flex flex-row justify-start w-full mb-5 mt-[-2vh]">
                                <input type="checkbox" name="savePass" className="m-0 w-5 h-5 border-2 border-[#2B4074]"/>
                                <label htmlFor="savePass" className="ml-1">Сохранить вход</label>
                                <a href="" className="ml-[50%] self-end text-[#2B4074]">Забыли пароль?</a>
                            </div>
                            <button className="mb-5 bg-[#2B4074] text-white h-[40px] w-[40%] self-center font-[600] rounded-[3px]">Вход</button>
                            <Link to="/register" className="mb-5 pt-1.5 border-[2px] border-[#2B4074] h-[40px] w-[40%] self-center font-[600] rounded-[3px] text-center">Регистрация</Link>
                        </form>
                    </div>
                </section>
            </section>

            <img src={HalfImage} alt="half_image" className="w-[50vw] h-[100vh]" />
        </section>
    );
};

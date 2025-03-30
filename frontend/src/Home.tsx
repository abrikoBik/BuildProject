import { Header } from "./Header";

export const Home = () => {
    return (
        <main>
            <Header />

            <section className="flex w-[100vw] h-[100vh] items-center justify-center text-2xl">
                <div className="align-center text-center">{`Добро пожаловать ${JSON.parse(localStorage.getItem('user')!).name}`}</div>
            </section>
        </main>
    );
};

import Sidebar from './Sidebar';
import Header from './Header';

const Projects = () => {
    const projects = Array.from({ length: 500 }, (_, i) => i + 1);

    return (
        <div className="flex flex-col h-screen">
            <Header />
            <div className="flex flex-row flex-1">
                <Sidebar show={true} choice={1} />
                <main className="w-[83%] ml-[17%] pr-5">
                    <h1 className="text-[#1D3266] text-4xl font-bold my-5">Мои проекты</h1>
                    <ul className="flex flex-row w-full gap-5 flex-wrap">
                        {projects.map((i) => (
                            <li key={i} className="flex flex-col w-[20em] h-[30em]">
                                <div
                                    style={{
                                        backgroundColor: `rgba(${Math.min(200 + i, 255)}, 0, ${Math.min(50 + i, 255)}, 1)`
                                    }}
                                    className="w-full h-[70%] border-3 rounded-xl flex items-center justify-center"
                                >
                                    image
                                </div>
                                <h3 className="w-full text-center text-xl my-2">{`Проект ${i}`}</h3>
                                <button
                                    className="w-full h-15 text-center border-2 bg-palette-color4 pt-1 rounded-xl font-semibold text-xl text-white self-center"
                                >
                                    Перейти к проекту
                                </button>
                            </li>
                        ))}
                    </ul>
                </main>
            </div>
        </div>
    );
};

export default Projects;

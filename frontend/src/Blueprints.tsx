import { useState, useRef, ChangeEvent } from 'react';
import axios from 'axios';
import Header from './Header';
import Sidebar from './Sidebar';
import Img1 from '../src/assets/image-1.png?url';
import Img2 from '../src/assets/image-2.png?url';
import Img3 from '../src/assets/image-3.png?url';
import Img4 from '../src/assets/image.png?url';

interface BlueprintProp {
    src: string,
    title: string,
    description?: string;
}

interface EditableCellProps {
    value: string;
    onSave?: (newValue: string) => void;
    onChange?: (newValue: string) => void;
    inputStyle?: string;
    textStyle?: string;
    textArea?: boolean;
}

const EditableCell = ({ value, onSave, onChange, inputStyle, textStyle, textArea = false }: EditableCellProps) => {
    const [isEdit, setEdit] = useState(false);
    const [curValue, setValue] = useState(value);

    const handleEditEnd = (event: React.KeyboardEvent<HTMLInputElement | HTMLTextAreaElement> | React.FocusEvent) => {
        if (event.type === 'keydown' && (event as React.KeyboardEvent).key !== 'Enter') {
            return;
        }

        setEdit(false);
        if (onSave) {
            onSave(curValue);
        }
    };

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setValue(e.target.value);

        if(onChange) {
            onChange(e.target.value);
        }
    };

    return (
    <div>
        {isEdit ? (
            textArea ? (
                <textarea
                    value={curValue}
                    onChange={handleChange}
                    onBlur={handleEditEnd}
                    onKeyDown={handleEditEnd}
                    autoFocus
                    className={inputStyle}
                    rows={4}
                />
            ) : (
                <input
                    type="text"
                    value={curValue}
                    onChange={handleChange}
                    onBlur={handleEditEnd}
                    onKeyDown={handleEditEnd}
                    autoFocus
                    className={inputStyle}
                />
            )
        ) : (
            <div 
                onDoubleClick={() => setEdit(true)}
                className={textStyle}
            >
                {curValue}
            </div>
        )}
    </div>
);
};

const Blueprints = () => {

    const fileInputRef = useRef<HTMLInputElement>(null);

    const [currentBlueprint, setCurrentBlueprint] = useState({
        src: "",
        title: "",
        description: ""
    });

    const [currentBlueprintIndex, setCurrentBlueprintIndex] = useState(-1);

    const [editVisible, setEditVisible] = useState(false);
    
    // Test description in future this should load from the server
    const [blueprintsArray, setBlueprintsArray] = useState<BlueprintProp[]>([
        {
            src: Img1,
            title: "План этажа",
            description: " Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc vestibulum urna blandit facilisis finibus. Curabitur faucibus suscipit tortor, et ultrices dui accumsan sed. Aenean risus sapien, suscipit at nisi ac, sagittis varius lacus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Sed gravida odio ac interdum eleifend. Sed condimentum neque id nunc scelerisque convallis. Ut sodales ultrices libero non cursus. Nulla in nibh euismod, congue lacus eu, consequat justo.Fusce sed ipsum vestibulum, posuere elit id, aliquam nulla. Duis eleifend elit sit amet enim condimentum feugiat non quis nibh. Nam purus arcu, pretium quis tincidunt id, fermentum in nunc. Donec porta elit eu odio lobortis, quis posuere nisl dapibus. Pellentesque viverra commodo tellus hendrerit pretium. Praesent et feugiat mauris, id fermentum massa. Proin ac lectus sit amet augue pulvinar cursus. Nullam hendrerit a lorem vel vestibulum. Morbi pharetra leo et erat maximus, a vulputate nulla ullamcorper. Vivamus quis nunc maximus, efficitur arcu rutrum, tincidunt ante. "
        },
        {
            src: Img2,
            title: "Каркас дома",
            description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc accumsan purus id sodales rutrum. Aenean tempor leo a tortor malesuada, at lobortis nulla elementum. Ut in eleifend nulla. Sed hendrerit interdum purus, et maximus enim luctus at. Aenean in tincidunt libero, ut rhoncus arcu. Cras sed nibh nisi. Etiam et risus est. Nulla placerat erat ut tincidunt pharetra. Praesent nisl turpis, interdum vel orci non, blandit pulvinar nisl.Fusce imperdiet accumsan neque, sit amet consectetur nisl. Aliquam vehicula lobortis dolor vitae finibus. Fusce sed posuere purus, sagittis imperdiet eros. Pellentesque orci ligula, porta ut ex non, vehicula sodales massa. Donec. ",
        },
        {
            src: Img3,
            title: "План фундамента",
            description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis porttitor metus dolor, ut fringilla turpis fringilla sit amet. Ut dictum, lectus a semper laoreet, urna purus viverra quam, eu tempor odio massa et dui. Phasellus laoreet lectus sit amet pretium fermentum. Integer sed sodales ligula. Quisque finibus euismod tortor id."
        },
        {
            src: Img4,
            title: "Чертеж фасада",
            description: "Фасад, общая высота 8м"
        },
        {
            src: Img1,
            title: "План этажа",
            description: "Первый этаж двери и мебель"
        },
        {
            src: Img2,
            title: "Каркас дома",
            description: "Схема строительной системы",
        },
        {
            src: Img3,
            title: "План фундамента",
            description: "Фундамент, вид сверху"
        },
        {
            src: Img4,
            title: "Чертеж фасада",
            description: "Фасад, общая высота 8м"
        },
        {
            src: Img1,
            title: "План этажа",
            description: "Первый этаж двери и мебель"
        },
        {
            src: Img2,
            title: "Каркас дома",
            description: "Схема строительной системы",
        },
        {
            src: Img3,
            title: "План фундамента",
            description: "Фундамент, вид сверху"
        },
        {
            src: Img4,
            title: "Чертеж фасада",
            description: "Фасад, общая высота 8м"
        }
    ]);

    const handleZoom = (e: React.MouseEvent<HTMLButtonElement>) => {
        const card = e.currentTarget.closest(".flex.flex-col.border-2") as HTMLElement;

        const imgElement = card?.querySelector("img.image") as HTMLImageElement;

        let scale : number = 1;
        const scrollSpeed : number = 0.1;
        const maxScale : number = 9;
        const minScale : number = 0.5;

        let startX : number;
        let startY : number;
        let curX : number = 0;
        let curY : number = 0;

        let isMouseDown : boolean = false;

        if (imgElement?.src) {
            console.log("Image src:", imgElement.src);
            
            const overlay = document.createElement('div');
            overlay.style.position = 'fixed';
            overlay.style.top = '0';
            overlay.style.left = '0';
            overlay.style.width = '100vw';
            overlay.style.height = '100vh';
            overlay.style.backgroundColor = 'rgba(0,0,0,0.4)';
            overlay.style.display = 'flex';
            overlay.style.justifyContent = 'center';
            overlay.style.alignItems = 'center';
            overlay.style.zIndex = '150';
            overlay.style.cursor = 'grab';

            const exitButton = document.createElement('a');
            exitButton.style.position = 'fixed';
            exitButton.style.display = 'flex';
            exitButton.style.justifyContent = 'center';
            exitButton.style.alignItems = 'center';
            exitButton.style.flex = '1';
            exitButton.style.top = '0';
            exitButton.style.right = '0';
            exitButton.style.width = '30px';
            exitButton.style.height = '30px';
            exitButton.style.backgroundColor = 'white';
            exitButton.style.borderRadius = '100%';
            exitButton.style.marginRight = '10px';
            exitButton.style.marginTop = '10px';
            exitButton.style.cursor = 'pointer';
            
            const exitIcon = document.createElement('span');
            exitIcon.className = 'material-symbols-outlined';
            exitIcon.textContent = 'close';
            exitIcon.style.userSelect = 'none';
            exitIcon.style.pointerEvents = 'none';

            const zoomedImg = new Image();
            zoomedImg.src = imgElement.src;
            zoomedImg.style.objectFit = 'contain';
            zoomedImg.style.width = '100px';
            zoomedImg.style.height = '100px';
            zoomedImg.style.userSelect = 'none';
            zoomedImg.style.pointerEvents = 'none';

            exitButton.appendChild(exitIcon);
            overlay.appendChild(zoomedImg);
            overlay.appendChild(exitButton);
            document.body.appendChild(overlay);
            document.body.style.overflow = 'hidden';

            exitButton.onclick = () => {
              document.body.removeChild(overlay);
              document.body.style.overflow = '';
            };

            overlay.addEventListener("wheel", function(event) {
                scale += event.deltaY > 0 ? -scrollSpeed : scrollSpeed;
                scale = Math.min(Math.max(scale, minScale), maxScale);

                zoomedImg.style.width = `${100 * scale}px`;
                zoomedImg.style.height = `${100 * scale}px`;
            });

            overlay.addEventListener("mousedown", function(event) {

                startX = event.clientX;
                startY = event.clientY;
                isMouseDown = true;
            });

            overlay.addEventListener("mouseup", function() {
                isMouseDown = false;
            });

            overlay.addEventListener("mousemove", function(event) {
                if(isMouseDown) {
                    curX += event.clientX - startX;
                    curY += event.clientY - startY;
                    startX = event.clientX;
                    startY = event.clientY;
                    zoomedImg.style.transform = `translate(${curX}px, ${curY}px)`;
                }
            });
        }
    };  

    const handleDownload = (e: React.MouseEvent<HTMLButtonElement>, title : string) => {
        const card = e.currentTarget.closest('.flex.flex-col.border-2') as HTMLElement;

        const imgElement = card?.querySelector('img.image') as HTMLImageElement;
        if(imgElement?.src)
        {
            const link = document.createElement("a");
            link.href = imgElement?.src;
            link.download = title;
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
        }
    };

    const handleEditEnter = (index : number) => {
        setCurrentBlueprint({
            src: blueprintsArray[index].src,
            title: blueprintsArray[index].title,
            description: blueprintsArray[index].description as string
        })
        setCurrentBlueprintIndex(index);
        setEditVisible(true);
    }

    const handleFileInput = () => {
        fileInputRef.current?.click();
    }

    const handleFileChange = async (event : ChangeEvent<HTMLInputElement>) => {
        const selectedFiles = (event.currentTarget as HTMLInputElement).files;
        if (selectedFiles) {
            for(let i = 0; i < selectedFiles.length; i++)
            {
                const postData = new FormData();
                postData.append("attachment", selectedFiles[i]);

                try {
                    const { data, status } = await axios.post(
                        '/api/attachments', 
                        postData, 
                        {
                            headers: {
                                'Content-Type': 'multipart/form-data'
                            }
                    });
                    if (status != 200) {
                        alert(data);
                        return;
                    }

                    alert("Файл успешно загружен!");
                    console.log(data);

                } catch (err : unknown) {
                    console.log(err);
                }

                console.log(`File name: ${selectedFiles[i].name}. File size: ${selectedFiles[i].size}. File type: ${selectedFiles[i].type}`);
            }
        }
    }

    return (
        <div className="flex flex-col h-screen">
            <Header />

            { editVisible ? 
                <section className="fixed left-0 top-0 w-screen h-screen bg-[rgba(0,0,0,0.4)] z-150">
                    <div className="flex flex-col relative left-1/2 top-1/2 -translate-1/2 w-[60%] h-[80%] rounded-xl border-2 border-palette-color4 bg-palette-color1">
                        <button onClick={() => setEditVisible(false)} type="button" className="hover:cursor-pointer hover:text-red-700 w-7 h-7 flex items-center justify-center self-end">
                            <span className="material-symbols-outlined">
                                close
                            </span>
                        </button>

                        <img src={currentBlueprint.src} alt={currentBlueprint.title}
                             className="h-60 border-y-2 border-palette-color4 object-contain bg-gray-100 mb-3" 
                        />
                        <button className="w-40 h-10 hover:cursor-pointer text-center bg-palette-color4 rounded-xl font-semibold text-white self-center">Поменять чертеж</button>

                        <div className="grid grid-rows-4 h-full min-h-0 overflow-auto p-5">
                            <div className="flex flex-col">
                                <label className="text-sm font-medium mb-1">Название</label>
                                <EditableCell 
                                    value={currentBlueprint.title} 
                                    inputStyle="w-full p-2 border rounded"
                                    textStyle="p-2 cursor-text hover:bg-gray-100 rounded"
                                    onSave={(newValue : string) => {
                                        const updated = [...blueprintsArray];
                                        updated[currentBlueprintIndex].title = newValue || '-';
                                        setBlueprintsArray(updated);
                                    }}
                                />
                            </div>
                            
                            <div className="flex flex-col min-h-0 h-full">
                                <label className="text-sm font-medium mb-1">Описание</label>
                                <div className="flex-1 min-h-0">
                                  <EditableCell 
                                    value={currentBlueprint.description} 
                                    inputStyle="w-full h-50 p-2 border rounded"
                                    textStyle="h-full p-2 cursor-text hover:bg-gray-100 rounded"
                                    textArea={true}
                                    onSave={(newValue : string) => {
                                        const updated = [...blueprintsArray];
                                        updated[currentBlueprintIndex].description = newValue || '-';
                                        setBlueprintsArray(updated);
                                    }}
                                  />
                                </div>
                            </div>
                        </div>
                    </div>
                </section> :
                <></>
            }

            <div className={`flex flex-row flex-1 min-h-0 ${editVisible ? "overflow-hidden" : "overflow-auto"}`}>
                <Sidebar show={true} choice={3}/>
                <main className="w-[83%] ml-[17%] pr-5">
                    <h1 className="text-[#1D3266] text-4xl font-bold mt-5 ">Чертежи и планы</h1>
                    <div className="flex flex-row gap-10 mb-5">
                        <h3 className="text-lg">Заказчик: Иванов И. С.</h3>
                        <h3 className="text-lg">Объект: Московкая д. 68</h3>
                    </div>

                    <section className="w-full h-full grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                        {/* Кнопка загрузки - занимает одну ячейку */}
                        <div>
                            <input 
                                type="file" 
                                ref={fileInputRef}
                                onChange={handleFileChange}
                                className="hidden absolute"
                                multiple
                                accept="image/png, image/jpeg, image/jpg"
                            />
                            <button
                                onClick={handleFileInput}
                                className="h-90 w-full flex bg-palette-color4 flex-col items-center justify-center rounded-xl font-semibold text-xl text-white hover:cursor-pointer"
                            >
                                <span className="material-symbols-outlined-huge text-4xl mb-2">upload</span>
                                <div>Загрузить файлы</div>
                            </button>
                        </div>

                        {/* Карточки чертежей */}
                        {blueprintsArray.map((item, index) => (
                            <div key={index} className="flex flex-col h-90 border-2 border-palette-color4 rounded-xl p-4 hover:shadow-md transition-shadow">
                                <img 
                                    src={item.src} 
                                    alt={item.title} 
                                    className="rounded-lg w-full h-40 object-contain bg-gray-100 image"
                                />
                                <section className="flex flex-col mt-4 gap-3 flex-grow">
                                    <div>
                                        <div className="text-lg font-bold mb-1">{item.title}</div>
                                        <div className="text-gray-600 text-sm">{item.description?.length as number < 100 ? item.description : item.description?.slice(0, 97) + "..."}</div>
                                    </div>

                                    <div className="flex flex-row gap-3 ml-auto mt-auto">
                                        <button onClick={() => handleEditEnter(index)} 
                                                className="w-8 h-8 hover:cursor-pointer rounded-lg bg-gray-300 transition-colors hover:bg-gray-400 flex items-center justify-center"
                                        >
                                            <span className="material-symbols-outlined text-base">edit</span>
                                        </button>
                                        <button onClick={(e: React.MouseEvent<HTMLButtonElement>) => handleDownload(e, item.title)} 
                                                className="w-8 h-8 hover:cursor-pointer select-none rounded-lg bg-gray-300 transition-colors hover:bg-gray-400 hover:bg-palette-color2 flex items-center justify-center"
                                        >
                                            <span className="material-symbols-outlined text-base">download</span>
                                        </button>
                                        <button 
                                            className="w-8 h-8 hover:cursor-pointer rounded-lg bg-gray-300 transition-colors hover:bg-gray-400 flex items-center justify-center"
                                            onClick={handleZoom}
                                        >
                                            <span className="material-symbols-outlined text-base">search</span>
                                        </button>
                                    </div>
                                </section>
                            </div>
                        ))}
                    </section>
                </main>
            </div>
        </div>
    );
}

export default Blueprints;

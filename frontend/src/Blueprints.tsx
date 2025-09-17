import { useState, useRef, ChangeEvent, useEffect } from 'react';
import { BlueprintProp } from './types/blueprint.ts';
import axios from 'axios';
import Header from './Header';
import Sidebar from './Sidebar';
import { useCookies } from 'react-cookie';

/*const EditableCell = ({ value, onSave, onChange, inputStyle, textStyle, textArea = false }: EditableCellProps) => {
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
};*/

const Blueprints = () => {

    const fileInputRef = useRef<HTMLInputElement>(null);
    const token = useCookies(["access-token"])[0]["access-token"];

    const [currentBlueprint] = useState({
        src: "",
        title: "",
        description: ""
    });

    //const [currentBlueprintIndex, setCurrentBlueprintIndex] = useState(-1);
    const [currentImages, setCurrentImages] = useState<string[]>([""]);

    const [editVisible, setEditVisible] = useState(false);
    
    // Test description in future this should load from the server
    const [blueprintsArray, setBlueprintsArray] = useState<BlueprintProp[]>([
        
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

    const getImages = async (blueprints: BlueprintProp[]) => {
    const imagePromises = blueprints.map(async (item, index) => {
        try {
            const { data } = await axios.get(
                `/api/projects/2/attachments/${item.uuid}`,
                {
                    headers: {
                        "Authorization": `Bearer ${token}`,
                    }
                }
            );

            let mimeType = 'image/png';
            if (data.charAt(0) === '/') {
                mimeType = 'image/jpeg';
            } else if (data.charAt(0) === 'i') {
                mimeType = 'image/png';
            }

            return {
                index,
                url: `data:${mimeType};base64,${data}`
            };
        } catch (err) {
            console.warn(`Ошибка загрузки изображения ${item.uuid}:`, err);
            return {
                index,
                url: ""
            };
        }
    });
    const results = await Promise.all(imagePromises);

    const orderedImages = Array(blueprints.length).fill("");
    results.forEach(result => {
        orderedImages[result.index] = result.url;
    });

    setCurrentImages(orderedImages);
};

    const loadBlueprints = async () => {
        try {
            const { data, status } = await axios.get(
                '/api/projects/2/attachments',
                {
                    headers: {
                        "Authorization": `Bearer ${token}`,
                    }
                }
            );
            console.info(data);

            setBlueprintsArray(data);

            if(status != 200)
            {
                alert(`${status}`);
            }

            await getImages(data);
        }
        catch (err : unknown) {
            console.warn(`${err}`);
        }
    }

    /*const handleEditEnter = (index : number) => {
        setCurrentBlueprint({
            src: blueprintsArray[index].src,
            title: blueprintsArray[index].title,
            description: blueprintsArray[index].description as string
        })
        setCurrentBlueprintIndex(index);
        setEditVisible(true);
    }*/

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
                        '/api/projects/2/attachments', 
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

    useEffect(() => {
        loadBlueprints();
    }, [])

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
                                {/*<EditableCell 
                                    value={currentBlueprint.title} 
                                    inputStyle="w-full p-2 border rounded"
                                    textStyle="p-2 cursor-text hover:bg-gray-100 rounded"
                                    onSave={(newValue : string) => {
                                        const updated = [...blueprintsArray];
                                        updated[currentBlueprintIndex].title = newValue || '-';
                                        setBlueprintsArray(updated);
                                    }}
                                />*/}
                            </div>
                            
                            <div className="flex flex-col min-h-0 h-full">
                                <label className="text-sm font-medium mb-1">Описание</label>
                                <div className="flex-1 min-h-0">
                                  {/*<EditableCell 
                                    value={currentBlueprint.description} 
                                    inputStyle="w-full h-50 p-2 border rounded"
                                    textStyle="h-full p-2 cursor-text hover:bg-gray-100 rounded"
                                    textArea={true}
                                    onSave={(newValue : string) => {
                                        const updated = [...blueprintsArray];
                                        updated[currentBlueprintIndex].description = newValue || '-';
                                        setBlueprintsArray(updated);
                                    }}
                                  />*/}
                                </div>
                            </div>
                        </div>
                    </div>
                </section> :
                <></>
            }

            <div className={`flex flex-row flex-1 min-h-0 ${editVisible ? "overflow-hidden" : "overflow-auto"} mt-20 `}>
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
                                    src={currentImages[index]} 
                                    alt={item.name} 
                                    className="rounded-lg w-full h-40 object-contain bg-gray-100 image"
                                />
                                <section className="flex flex-col mt-4 gap-3 flex-grow">
                                    <div>
                                        <div className="text-lg font-bold mb-1">{item.name}</div>
                                        {/*<div className="text-gray-600 text-sm">{item.description?.length as number < 100 ? item.description : item.description?.slice(0, 97) + "..."}</div>*/}
                                    </div>

                                    <div className="flex flex-row gap-3 ml-auto mt-auto">
                                        <button 
                                                className="w-8 h-8 hover:cursor-pointer rounded-lg bg-gray-300 transition-colors hover:bg-gray-400 flex items-center justify-center"
                                        >
                                            <span className="material-symbols-outlined text-base">edit</span>
                                        </button>
                                        <button onClick={(e: React.MouseEvent<HTMLButtonElement>) => handleDownload(e, item.name)} 
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

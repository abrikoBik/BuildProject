import { createRoot } from "react-dom/client";
import { StrictMode } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router";
import { CookiesProvider } from "react-cookie";
import "./index.css";
import Home from "./Home.tsx";
import Login from "./Login.tsx";
import Register from "./Register.tsx";
import Estimates from "./Estimates.tsx";
import Projects from "./Projects.tsx";
import Blueprints from "./Blueprints.tsx";

{/* DEV */}
import GanttChart from "./Gantt.tsx";

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <CookiesProvider>
            <Router>
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                    <Route path="/estimates" element={<Estimates />} />
                    <Route path="/blueprints" element={<Blueprints />} />
                    <Route path="/projects" element={<Projects />}/>
                    <Route path="/ganttDEV" element={<GanttChart />}/>
                </Routes>
            </Router>
        </CookiesProvider>
    </StrictMode>,
);

import { createRoot } from "react-dom/client";
import { StrictMode } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router";
import "./index.css";
import { Home } from "./Home.tsx";
import { Login } from "./Login.tsx";
import { Register } from "./Register.tsx";

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <Router>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/login" element={<Login />} />
                <Route path="/register" element={<Register />} />
            </Routes>
        </Router>
    </StrictMode>,
);

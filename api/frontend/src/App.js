import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import LoginPage from "./components/LoginPage";
import DashboardAdmin from "./pages/DashboardAdmin";
import DashboardSocio from "./pages/DashboardSocio";
import './App.css';

function App() {
  return (
    <BrowserRouter>
      <div className="app-container">
        <Routes>
          <Route path="/" element={<LoginPage />} />
          <Route path="/admin" element={<DashboardAdmin />} />
          <Route path="/socio" element={<DashboardSocio />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;

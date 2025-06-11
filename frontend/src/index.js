// index.js
import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import LoginPage from './pages/LoginPage.jsx';
import DashboardAdmin from './pages/DashboardAdmin.jsx';
import DashboardSocio from './pages/DashboardSocio.jsx';
import ActividadDetalle from './pages/ActividadDetalle.jsx';
import './index.css';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<LoginPage />} />
      <Route path="/admin" element={<DashboardAdmin />} />
      <Route path="/socio" element={<DashboardSocio />} />
      <Route path="/actividad/:id" element={<ActividadDetalle />} />
    </Routes>
  </BrowserRouter>
);
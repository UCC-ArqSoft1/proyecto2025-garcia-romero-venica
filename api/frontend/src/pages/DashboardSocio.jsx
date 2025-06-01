import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import './DashboardSocio.css';

const actividadesFake = [
  { id: 1, titulo: "Fútbol", horario: "18:00 a 20:00", profesor: "Profesor: Carlos Pérez", descripcion: "Partido amistoso de fútbol 5" },
  { id: 2, titulo: "Yoga", horario: "09:00 a 11:00", profesor: "Profesor: Laura Gómez", descripcion: "Clase de yoga para principiantes" },
  { id: 3, titulo: "Crossfit", horario: "20:00 a 22:00", profesor: "Profesor: Martín Díaz", descripcion: "Entrenamiento de alta intensidad" },
];

function DashboardSocio() {
  const [search, setSearch] = useState("");
  const [actividades] = useState(actividadesFake);
  const navigate = useNavigate();

  const actividadesFiltradas = actividades.filter((act) =>
    act.titulo.toLowerCase().includes(search.toLowerCase()) ||
    act.horario.includes(search) ||
    act.profesor.toLowerCase().includes(search.toLowerCase())
  );

  const handleInscribirse = (actividad) => {
    Swal.fire({
      title: '¡Inscripción exitosa!',
      text: `Te has inscrito a: ${actividad.titulo}`,
      icon: 'success',
      confirmButtonColor: '#3085d6',
      confirmButtonText: 'OK'
    });
  };

  const handleVerDetalle = (id) => {
    navigate(`/actividad/${id}`);
  };

  return (
    <main className="socio-container">
      <h1 className="socio-titulo">Actividades Disponibles</h1>

      <input
        type="text"
        placeholder="Buscar por título, horario o profesor"
        value={search}
        onChange={(e) => setSearch(e.target.value)}
        className="socio-input"
      />

      <ul className="socio-lista">
        {actividadesFiltradas.length === 0 && <li>No se encontraron actividades</li>}
        {actividadesFiltradas.map((act) => (
          <li key={act.id} className="socio-itemLista">
            <div>
              <strong>{act.titulo}</strong> - {act.horario} - {act.profesor}<br />
              <span>{act.descripcion}</span>
            </div>
            <div style={{ marginTop: "10px", display: "flex", gap: "10px" }}>
              <button onClick={() => handleInscribirse(act)} className="socio-btn">Inscribirse</button>
              <button onClick={() => handleVerDetalle(act.id)} className="socio-btn">Información</button>
            </div>
          </li>
        ))}
      </ul>
    </main>
  );
}

export default DashboardSocio;

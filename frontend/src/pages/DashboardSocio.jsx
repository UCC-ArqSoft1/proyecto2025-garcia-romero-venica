import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import "./DashboardSocio.css";

const actividadesFake = [
  { id: 1, titulo: "Fútbol", horario: "18:00 a 20:00", profesor: "Carlos Pérez", descripcion: "Partido amistoso de fútbol 5" },
  { id: 2, titulo: "Yoga", horario: "09:00 a 11:00", profesor: "Laura Gómez", descripcion: "Clase de yoga para principiantes" },
  { id: 3, titulo: "Crossfit", horario: "20:00 a 22:00", profesor: "Martín Díaz", descripcion: "Entrenamiento de alta intensidad" },
];

function DashboardSocio() {
  const [search, setSearch] = useState("");
  const [actividades] = useState(actividadesFake);
  const [inscripciones, setInscripciones] = useState([]);
  const navigate = useNavigate();

  const actividadesFiltradas = actividades.filter((act) =>
    act.titulo.toLowerCase().includes(search.toLowerCase()) ||
    act.horario.includes(search) ||
    act.profesor.toLowerCase().includes(search.toLowerCase())
  );

  const handleInscribirse = (actividad) => {
    if (!inscripciones.find((a) => a.id === actividad.id)) {
      setInscripciones([...inscripciones, actividad]);
      Swal.fire("¡Inscripción exitosa!", `Te has inscrito a: ${actividad.titulo}`, "success");
    } else {
      Swal.fire("Ya estás inscrito", `Ya estás inscrito a: ${actividad.titulo}`, "info");
    }
  };

  const handleDesinscribirse = (id) => {
    setInscripciones(inscripciones.filter((a) => a.id !== id));
    Swal.fire("Cancelado", "Te has desinscripto de la actividad", "info");
  };

  const handleVerDetalle = (id) => {
    navigate(`/actividad/${id}`);
  };

  return (
    <main className="socio-container">
      <button onClick={() => navigate(-1)} className="volver-btn">← Volver</button>
      <h1 className="socio-titulo">Actividades Disponibles</h1>

      <input
        type="text"
        placeholder="Buscar por título, horario o profesor"
        value={search}
        onChange={(e) => setSearch(e.target.value)}
        className="socio-input"
      />

      <div className="socio-grid">
        <section className="socio-card">
          <h2>Buscar e Inscribirse</h2>
          <ul className="socio-lista">
            {actividadesFiltradas.length === 0 && <li>No se encontraron actividades</li>}
            {actividadesFiltradas.map((act) => (
              <li key={act.id} className="socio-item">
                <div>
                  <strong>{act.titulo}</strong> - {act.horario} - Profesor: {act.profesor}
                  <p>{act.descripcion}</p>
                </div>
                <div className="socio-botones">
                  <button onClick={() => handleInscribirse(act)} className="socio-btn">Inscribirse</button>
                  <button onClick={() => handleVerDetalle(act.id)} className="socio-btn info">Información</button>
                </div>
              </li>
            ))}
          </ul>
        </section>

        <section className="socio-card">
          <h2>Actividades Inscritas</h2>
          {inscripciones.length === 0 ? (
            <p className="socio-info">Todavía no te has inscrito a ninguna actividad.</p>
          ) : (
            <ul className="socio-lista">
              {inscripciones.map((act) => (
                <li key={act.id} className="socio-item inscrito">
                  <strong>{act.titulo}</strong> - {act.horario}<br />
                  <small>Profesor: {act.profesor}</small>
                  <div className="socio-botones">
                    <button onClick={() => handleDesinscribirse(act.id)} className="socio-btn cancelar">Desinscribirse</button>
                  </div>
                </li>
              ))}
            </ul>
          )}
        </section>
      </div>
    </main>
  );
}

export default DashboardSocio;

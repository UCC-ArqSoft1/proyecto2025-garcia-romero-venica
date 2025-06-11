// DashboardSocio.jsx
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import "./DashboardSocio.css";

function DashboardSocio() {
  const [search, setSearch] = useState("");
  const [actividades, setActividades] = useState([]);
  const [inscripciones, setInscripciones] = useState([]);
  const navigate = useNavigate();
  const token = localStorage.getItem("token");

  useEffect(() => {
    fetch("http://localhost:8080/actividades")
      .then(res => res.json())
      .then(data => setActividades(data))
      .catch(err => console.error("Error al cargar actividades", err));

    fetch("http://localhost:8080/api/inscripciones/usuario", {
      headers: { Authorization: `Bearer ${token}` }
    })
      .then(res => res.json())
      .then(data => setInscripciones(data))
      .catch(err => console.error("Error al cargar inscripciones", err));
  }, [token]);

  const actividadesFiltradas = actividades.filter((act) =>
    act.titulo.toLowerCase().includes(search.toLowerCase()) ||
    act.horario.includes(search) ||
    act.profesor.toLowerCase().includes(search.toLowerCase())
  );

  const handleInscribirse = (actividad) => {
    fetch("http://localhost:8080/api/inscripciones", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`
      },
      body: JSON.stringify({ actividad_id: actividad.id })
    })
      .then(res => res.json())
      .then(data => {
        Swal.fire("¡Inscripción exitosa!", data.mensaje, "success");
        setInscripciones(prev => [...prev, actividad]);
      })
      .catch(err => {
        Swal.fire("Error", "No se pudo realizar la inscripción", "error");
        console.error(err);
      });
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

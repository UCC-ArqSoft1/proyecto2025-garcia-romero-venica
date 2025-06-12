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
  const userID = localStorage.getItem("userID");

  useEffect(() => {
    fetch("http://localhost:8080/actividades")
      .then(res => res.json())
      .then(data => setActividades(data))
      .catch(err => console.error("Error al cargar actividades", err));

    fetch(`http://localhost:8080/inscripciones/${userID}`, {
      headers: { Authorization: `Bearer ${token}` }
    })
      .then(res => res.json())
      .then(data => {
        if (Array.isArray(data)) {
          setInscripciones(data);
        } else {
          console.warn("Respuesta inesperada:", data);
          setInscripciones([]);
        }
      })
      .catch(err => {
        console.error("Error al cargar inscripciones:", err);
        setInscripciones([]);
      });
  }, [token, userID]);

  const actividadesFiltradas = actividades.filter((act) =>
    (act?.nombre || "").toLowerCase().includes(search.toLowerCase()) ||
    (act?.horario || "").toLowerCase().includes(search.toLowerCase()) ||
    (act?.profesor || "").toLowerCase().includes(search.toLowerCase())
  );

  const handleInscribirse = (actividad) => {
    fetch("http://localhost:8080/inscripciones", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`
      },
      body: JSON.stringify({ actividad_id: actividad.id, usuario_id: parseInt(userID) })
    })
      .then(res => res.json())
      .then(data => {
        Swal.fire("¡Inscripción exitosa!", "Te has inscripto correctamente.", "success");
        setInscripciones(prev => [...prev, actividad]);
      })
      .catch(err => {
        Swal.fire("Error", "No se pudo realizar la inscripción", "error");
        console.error(err);
      });
  };

  const handleCancelarInscripcion = (actividad) => {
    Swal.fire({
      title: "¿Cancelar inscripción?",
      text: `¿Deseás cancelar tu inscripción a "${actividad.nombre}"?`,
      icon: "warning",
      showCancelButton: true,
      confirmButtonText: "Sí, cancelar",
      cancelButtonText: "No"
    }).then(async (result) => {
      if (result.isConfirmed) {
        try {
          const res = await fetch(`http://localhost:8080/inscripciones/${actividad.id}`, {
            method: "DELETE",
            headers: { Authorization: `Bearer ${token}` }
          });

          if (!res.ok) {
            const data = await res.json();
            throw new Error(data.error || "No se pudo cancelar");
          }

          Swal.fire("Cancelado", "Te has desinscripto correctamente.", "success");
          setInscripciones(inscripciones.filter(i => i.id !== actividad.id));
        } catch (err) {
          Swal.fire("Error", err.message, "error");
        }
      }
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
        placeholder="Buscar por nombre, horario o profesor"
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
                  <strong>{act.nombre}</strong> - {act.horario} - Profesor: {act.profesor}
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
                  <strong>{act.nombre}</strong> - {act.horario}<br />
                  <small>Profesor: {act.profesor}</small>
                  <div className="socio-botones">
                    <button onClick={() => handleCancelarInscripcion(act)} className="socio-btn cancelar">
                      Cancelar inscripción
                    </button>
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

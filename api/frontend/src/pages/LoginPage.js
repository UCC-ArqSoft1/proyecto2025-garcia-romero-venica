import React, { useState } from "react";
import "./LoginPage.css";

function LoginPage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = (e) => {
    e.preventDefault();
    // Acá iría la lógica para enviar al backend (con fetch o axios)
    alert(`Login enviado: ${username}`);
    };

    return (
    <div className="login-container">
        <h2>Iniciar Sesión</h2>
        <form onSubmit={handleSubmit}>
            <input
            type="text"
            placeholder="Usuario"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
            />
            <input
            type="password"
            placeholder="Contraseña"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
        />
        <button type="submit">Ingresar</button>
        </form>
    </div>
    );
}

export default LoginPage;

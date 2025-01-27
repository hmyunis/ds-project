import React, { useState } from "react";
import axios from "axios";
import "./SignupLogin.css";

interface Props {
  onAuthSuccess: (username: string) => void;
}

const SignupLogin: React.FC<Props> = ({ onAuthSuccess }) => {
  const [authMode, setAuthMode] = useState<"signup" | "login">("login");
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [error, setError] = useState<string>("");

  const handleAuth = async (e: React.FormEvent) => {
    e.preventDefault();
    const endpoint = authMode === "signup" ? "/signup" : "/login";

    try {
      const response = await axios.post(`http://localhost:8080${endpoint}`, {
        username,
        password,
      });

      if (authMode === "login") {
        onAuthSuccess(username); // Notify parent of successful login
      } else {
        alert("Signup successful! Please log in.");
        setAuthMode("login");
      }
      setError("");
    } catch (err: any) {
      setError(err.response?.data || "An error occurred");
    }
  };

  return (
    <div className="auth-container">
      <h2>{authMode === "signup" ? "Sign Up" : "Log In"}</h2>
      <form onSubmit={handleAuth}>
        <div>
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>
        <div>
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit">{authMode === "signup" ? "Sign Up" : "Log In"}</button>
      </form>
      <p>
        {authMode === "signup" ? "Already have an account?" : "Don't have an account?"}{" "}
        <button onClick={() => setAuthMode(authMode === "signup" ? "login" : "signup")}>
          {authMode === "signup" ? "Log In" : "Sign Up"}
        </button>
      </p>
      {error && <p className="error">{error}</p>}
    </div>
  );
};

export default SignupLogin;

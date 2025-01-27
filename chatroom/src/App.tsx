import React, { useState } from "react";
import SignupLogin from "./SignupLogin";
import ChatRoom from "./ChatRoom";

const App: React.FC = () => {
  const [authenticated, setAuthenticated] = useState<boolean>(false);
  const [username, setUsername] = useState<string>("");

  const handleAuthSuccess = (username: string) => {
    setAuthenticated(true);
    setUsername(username);
  };

  return (
    <div className="App">
      {!authenticated ? (
        <SignupLogin onAuthSuccess={handleAuthSuccess} />
      ) : (
        <ChatRoom username={username} />
      )}
    </div>
  );
};

export default App;

import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import TestComponent from "./components/test";
import Redux from "./components/redux";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path={`/`} element={<TestComponent />} />
        <Route path={`/redux`} element={<Redux />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;

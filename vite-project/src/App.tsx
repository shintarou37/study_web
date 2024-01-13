import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import TestComponent from "./components/test";
import Redux from "./components/redux";
import Form from "./components/form";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path={`/`} element={<TestComponent />} />
        <Route path={`/redux`} element={<Redux />} />
        <Route path={`/form`} element={<Form />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;

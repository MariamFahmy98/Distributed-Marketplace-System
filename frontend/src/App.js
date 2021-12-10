import React from "react";
import { Routes, Route } from "react-router-dom";

import MainNav from "./components/MainNav";

import Home from "./pages/Home";
import Profile from "./pages/Profile";
import Stores from "./pages/Stores";
import Store from "./pages/Store";
import ProductPage from "./pages/Product";

function App() {
  return (
    <MainNav>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/profile" element={<Profile />} />
        <Route path="/store" element={<Stores />} />
        <Route path="/store/:id" element={<Store />} />
        <Route path="/product/:id" element={<ProductPage />} />
      </Routes>
    </MainNav>
  );
}

export default App;

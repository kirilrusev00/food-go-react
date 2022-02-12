import React from 'react';
import { CssBaseline } from '@material-ui/core';
import {
  BrowserRouter, Route, Navigate, Routes,
} from 'react-router-dom';
import Footer from './components/footer';
import Header from './components/header';
import MainPage from './pages/main-page';
import SearchPage from './pages/search-page';

export default function App() {
  return (
    <BrowserRouter>
      <CssBaseline />
      <Header />
      <Routes>
        <Route path="/search" element={<SearchPage />} />
        <Route path="/" element={<MainPage />} />
        <Route path="*" element={<Navigate to="/" />} />
        {' '}
      </Routes>
      <Footer />
    </BrowserRouter>
  );
}

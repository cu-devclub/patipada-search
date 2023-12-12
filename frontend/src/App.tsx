import { BrowserRouter, Routes, Route } from "react-router-dom";
import Wrapper from "./pages/Wrapper";
import { AdminLoginPage, AdminResetPasswordPage } from "./pages/Admin";
export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Wrapper />} />
        <Route path="/admin" element={<AdminLoginPage />} />
        <Route path="/admin-reset-password" element={<AdminResetPasswordPage />} />
      </Routes>
    </BrowserRouter>
  );
}

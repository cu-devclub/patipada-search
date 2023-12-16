import { BrowserRouter, Routes, Route } from "react-router-dom";
import { SearchMiddleware } from "./pages/Search";
import { AdminResetPasswordPage, AdminMiddleware } from "./pages/Admin";
export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<SearchMiddleware />} />
        <Route path="/admin" element={<AdminMiddleware />} />
        <Route
          path="/admin-reset-password"
          element={<AdminResetPasswordPage />}
        />
      </Routes>
    </BrowserRouter>
  );
}

import { BrowserRouter, Routes, Route } from "react-router-dom";
import { SearchMiddleware } from "./pages/Search";
import { ResetPasswordPage, UserMiddleware,LoginPage,RegisterPage } from "./pages/User";
export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<SearchMiddleware />} />
        <Route path="/user/">
          <Route path="" element={<UserMiddleware />} />
          <Route path="login" element={<LoginPage />} />
          <Route path="reset-password" element={<ResetPasswordPage />} />
          <Route path="register" element={<RegisterPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

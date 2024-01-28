import { AuthorizeAdmin } from "../../functions";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Role } from "../../constant";
import { AdminBasePage } from "./AdminBasePage";
function AdminDashboard() {
  const navigate = useNavigate();
  useEffect(() => {
    (async () => {
      const isAuthorize = await AuthorizeAdmin(Role.ADMIN);
      if (isAuthorize === false) {
        navigate("/user/login");
      }
    })();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <AdminBasePage activePage="Dashboard">
      <></>
    </AdminBasePage>
  );
}

export default AdminDashboard;

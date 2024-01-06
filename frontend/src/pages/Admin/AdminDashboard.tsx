import { BaseHeader } from "../../components";
import { Heading } from "@chakra-ui/react";
import { AuthorizeAdmin } from "../../functions";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Role } from "../../constant";
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
    <BaseHeader>
      <Heading>Dashboard</Heading>
    </BaseHeader>
  );
}

export default AdminDashboard;

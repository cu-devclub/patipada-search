import { getCookie } from "typescript-cookie";
import { MessageToast } from "../../components";
import { ToastStatus } from "../../constant";
import { AdminLoginPage, DataManagementPage } from ".";
function AdminMiddleware() {
  const { addToast } = MessageToast();

  const token = getCookie("token");
  if (token) {
    return <DataManagementPage />;
  } else {
    addToast({
      description: "Please login first",
      status: ToastStatus.WARNING,
    });
    return <AdminLoginPage />;
  }
}

export default AdminMiddleware;

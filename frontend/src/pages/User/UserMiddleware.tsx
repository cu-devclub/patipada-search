import { getCookie } from "typescript-cookie";
import { MessageToast } from "../../components";
import { ToastStatus } from "../../constant";
import { LoginPage, DataManagementPage } from ".";

/**
 * A middleware function that checks if the user is an admin. If the user is an admin and has a valid token,
 * the function renders the DataManagementPage component. If the user is not an admin or does not have a valid token,
 * the function renders the LoginPage component and displays a warning toast.
 *
 * @return {JSX.Element} The rendered component based on the user's admin status and token validity.
 */
function UserMiddleware() {
  const { addToast } = MessageToast();

  const token = getCookie("token");
  if (token) {
    return <DataManagementPage />;
  } else {
    addToast({
      description: "Please login first",
      status: ToastStatus.WARNING,
    });
    return <LoginPage />;
  }
}

export default UserMiddleware;

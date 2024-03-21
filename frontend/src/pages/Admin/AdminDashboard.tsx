import { AuthorizeAdmin } from "../../functions";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Role, ToastStatus } from "../../constant";
import { AdminBasePage } from "./AdminBasePage";
import { StatGroup } from "@chakra-ui/react";
import { RecordStat, RequestStat, Userstat } from "../../components/stat";
import { AuthSummary, User } from "../../models/user";
import { getAllUsers } from "../../service/user/getUser";
import { MessageToast } from "../../components";
import { getDataSummary } from "../../service/data/getRequest";
import { RecordSummary } from "../../models/qa";
import { RequestSummary } from "../../models/request";
function AdminDashboard() {
  const navigate = useNavigate();
  const { addToast } = MessageToast();
  const [authSummary, setAuthSummary] = useState<AuthSummary>({
    sumTotal: 0,
    totalUser: 0,
    totalAdmin: 0,
    totalSuperAdmin: 0,
  });

  const [recordSummary, setRecordSummary] = useState<RecordSummary>({
    recordAmount: 0,
    youtubeClipAmount: 0,
  });

  const [requestSummary, setRequestSummary] = useState<RequestSummary>({
    requestAmount: 0,
    reviewedAmount: 0,
    pendingAmount: 0,
  });

  useEffect(() => {
    (async () => {
      const isAuthorize = await AuthorizeAdmin(Role.ADMIN);
      if (isAuthorize === false) {
        navigate("/user/login");
      }
    })();

    const getUsers = async () => {
      await getAllUsers()
        .then((res: User[]) => {
          const totalUser = res.filter(
            (user) => user.role === Role.USER
          ).length;
          const totalAdmin = res.filter(
            (user) => user.role === Role.ADMIN
          ).length;
          const totalSuperAdmin = res.filter(
            (user) => user.role === Role.SUPER_ADMIN
          ).length;
          setAuthSummary({
            sumTotal: res.length,
            totalUser,
            totalAdmin,
            totalSuperAdmin,
          });
        })
        .catch(() => {
          addToast({
            description: "เกิดข้อผิดพลาดขณะทำการดึงข้อมูล",
            status: ToastStatus.ERROR,
          });
        });
    };
    getUsers();

    const getDataSum = async () => {
      await getDataSummary()
        .then((res) => {
          if (res.recordSummary) {
            const recSum: RecordSummary = {
              recordAmount: res.recordSummary.recordAmount,
              youtubeClipAmount: res.recordSummary.youtubeClipAmount,
            };
            setRecordSummary(recSum);
          }

          if (res.requestSummary) {
            const reqSum: RequestSummary = {
              requestAmount: res.requestSummary.requestAmount,
              reviewedAmount: res.requestSummary.reviewedAmount,
              pendingAmount: res.requestSummary.pendingAmount,
            };
            setRequestSummary(reqSum);
          }
        })
        .catch(() => {
          addToast({
            description: "เกิดข้อผิดพลาดขณะทำการดึงข้อมูล",
            status: ToastStatus.ERROR,
          });
        });
    };
    getDataSum();

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <AdminBasePage activePage="Dashboard">
      <StatGroup pt={8}>
        <Userstat authSummary={authSummary} />
        <RecordStat recordSummary={recordSummary} />
        <RequestStat requestSummary={requestSummary} />
      </StatGroup>
    </AdminBasePage>
  );
}

export default AdminDashboard;

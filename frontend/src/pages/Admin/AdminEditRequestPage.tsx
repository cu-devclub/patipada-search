import { useParams } from "react-router-dom";
import {
  getRequestByParamsService,
  updateRequestService,
  insertRequestService,
} from "../../service/data";
import { useEffect, useState } from "react";
import { EditRequestForm } from "../../components/admin";
import {  MessageToast } from "../../components";
import { Request, createEncodeRequest, mapRequestToInsertRequestModels } from "../../models/request";
import { ToastStatus } from "../../constant";
import { useNavigate } from "react-router-dom";
import { getCookie } from "typescript-cookie";
import { AdminBasePage } from "./AdminBasePage";
function AdminEditRequestPage() {
  const { requestID } = useParams();
  const [data, setData] = useState<Request>();
  const { addToast } = MessageToast();
  const navigate = useNavigate();

  useEffect(() => {
    // 1. Try getting data from request first
    // 2. If not found, try searching
    const getRecord = async (requestID: string) => {
      await getRequestByParamsService({ requestID: requestID })
        .then((res) => {
          setData(res[0]);
          addToast({
            description: "ดึงข้อมูลสำเร็จ",
            status: ToastStatus.SUCCESS,
          });
        })
        .catch(() => {
          addToast({
            description: "เกิดข้อผิดพลาดขณะทำการดึงข้อมูล",
            status: ToastStatus.ERROR,
          });
        });
    };
    if (requestID) getRecord(requestID);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [requestID]);

  const submit = async (data: Request) => {
    if (data.status == "reviewed") {
      await insert(data);
      return;
    }

    data.approvedBy = getCookie("username") || "";
    data = createEncodeRequest(data);

    await updateRequestService(data)
      .then(() => {
        addToast({
          description: "อัพเดทคำขอแก้ไขสำเร็จ",
          status: ToastStatus.SUCCESS,
        });
        navigate(-1);
      })
      .catch(() => {
        addToast({
          description: "เกิดข้อผิดพลาดขณะอัพเดทคำขอแก้ไข",
          status: ToastStatus.ERROR,
        });
      });
  };

  const insert = async (data: Request) => {
    data.by = getCookie("username") || "";
    const insertRequestData = mapRequestToInsertRequestModels(data);
    await insertRequestService(insertRequestData)
      .then(() => {
        addToast({
          description: "ส่งคำขอแก้ไขสำเร็จ",
          status: ToastStatus.SUCCESS,
        });
      })
      .catch(() => {
        addToast({
          description: "เกิดข้อผิดพลาดขณะทำการส่งคำขอแก้ไข",
          status: ToastStatus.ERROR,
        });
      })
      .finally(() => {
        navigate(-1);
      });
  };

  return (
    <AdminBasePage activePage="Request" requestID={requestID}>
      {/* TODO: Create a section that show what requests is part of this request (older then this request) and allow user to redirect to that request  */}
      {data && <EditRequestForm data={data} submit={submit} />}
    </AdminBasePage>
  );
}

export default AdminEditRequestPage;
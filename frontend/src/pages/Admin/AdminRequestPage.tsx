import { AuthorizeAdmin, convertStatusWord, extractStringFromHTML } from "../../functions";
import { useEffect,useState } from "react";
import { useNavigate } from "react-router-dom";
import { Role, ToastStatus } from "../../constant";
import { AdminBasePage } from "./AdminBasePage";
import { RequestTable } from "../../components/admin";
import { getRequestByParams } from "../../service/data";
import { MessageToast } from "../../components";
import { Request } from "../../models/request";
function AdminRequestPage() {
  const navigate = useNavigate();
  const { addToast} = MessageToast();

  const [data,setdata] = useState<Request[]>([])


  useEffect(() => {
    (async () => {
      const isAuthorize = await AuthorizeAdmin(Role.ADMIN);
      if (isAuthorize === false) {
        navigate("/user/login");
      }
    })();

    const getRequest = async() => {
        await getRequestByParams({status:"pending"})
        .then((res) => {
          const r = res.map((item) => {
            item.startTime = extractStringFromHTML(item.startTime)
            item.endTime = extractStringFromHTML(item.endTime)
            item.question = extractStringFromHTML(item.question)
            item.answer = extractStringFromHTML(item.answer)
            item.createdAt = new Date(item.createdAt).toLocaleString("th-TH")
            item.status = convertStatusWord(item.status)
            return item
          })
  
          setdata(r)
          addToast({
            description: "ดึงข้อมูลสำเร็จ",
            status: ToastStatus.SUCCESS,
          });
        })
        .catch(() => {
          addToast({
            description: "เกิดข้อผิดพลาดขณะทำการดึงข้อมูล",
            status: ToastStatus.ERROR,
          })
        })
      }
      getRequest()
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <AdminBasePage activePage="Request">
      <RequestTable data={data} />
    </AdminBasePage>
  );
}

export default AdminRequestPage;

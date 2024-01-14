import { useParams } from "react-router-dom";
import { search } from "../../service/search";
import { getRequestByRecordIndex, insertRequest } from "../../service/data";
import { useEffect, useState } from "react";
import {
  EditRecordHeader,
  EditRecordForm,
} from "../../components/contributor/edit-record";
import { Grid } from "@chakra-ui/react";
import { Footer, MessageToast } from "../../components";
import {
  Request,
  mapDataItemToRequest,
  mapRequestToInsertRequestModels,
} from "../../models/request";
import { ToastStatus } from "../../constant";
import { useNavigate } from "react-router-dom";
import { getCookie } from "typescript-cookie";
function EditRecordPage() {
  const { recordID } = useParams();
  const [data, setData] = useState<Request>();
  const { addToast } = MessageToast();
  const navigate = useNavigate();

  useEffect(() => {
    // 1. Try getting data from request first
    // 2. If not found, try searching
    const getRecord = async (recordID: string) => {
      await getRequestByRecordIndex(recordID)
        .then(async (res) => {
          if (res.requestID == "NOT FOUND") {
            try {
              // Perform search when status is 404
              const searchResult = await search(recordID);
              if (searchResult.data.length === 0) {
                navigate("404");
              }
              const request = mapDataItemToRequest(searchResult.data[0]);
              setData(request);
              addToast({
                description: "ดึงข้อมูลสำเร็จ",
                status: ToastStatus.SUCCESS,
              });
            } catch (searchErr) {
              addToast({
                description: "เกิดข้อผิดพลาดขณะทำการดึงข้อมูล",
                status: ToastStatus.ERROR,
              });
            }
          } else {
            setData(res);
            addToast({
              description: "ดึงข้อมูลสำเร็จ",
              status: ToastStatus.SUCCESS,
            });
          }
        })
        .catch(() => {
          addToast({
            description: "เกิดข้อผิดพลาดขณะทำการดึงข้อมูล",
            status: ToastStatus.ERROR,
          });
        });
    };
    if (recordID) getRecord(recordID);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [recordID]);

  const submit = async (data: Request) => {
    data.by = getCookie("username") || "";
    const insertRequestData = mapRequestToInsertRequestModels(data);
    await insertRequest(insertRequestData)
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
    <Grid templateRows="auto 1fr auto" gap={4} w="full" minH="100svh">
      <EditRecordHeader />
      {data && <EditRecordForm data={data} submit={submit} />}
      <Footer />
    </Grid>
  );
}

export default EditRecordPage;

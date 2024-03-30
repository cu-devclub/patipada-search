import {
  AuthorizeAdmin,
  convertStatusWord,
  extractStringFromHTML,
} from "../../functions";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Role, ToastStatus } from "../../constant";
import { AdminBasePage } from "./AdminBasePage";
import { RequestTable } from "../../components/admin";
import {
  getRequestByParamsService,
  syncRequestRecordService,
  syncAllRequestRecordService,
} from "../../service/data";
import { MessageToast } from "../../components";
import { Request } from "../../models/request";
import { Flex, Button, useDisclosure } from "@chakra-ui/react";
import { ActionModal } from "../../components/modal";
function AdminRequestPage() {
  const navigate = useNavigate();
  const { addToast } = MessageToast();
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [modalTitle, setmodalTitle] = useState<string>("");
  const [modalBody, setmodalBody] = useState<string>("");
  const [modalConfirm, setmodalConfirm] = useState<() => void>(() => () => {});

  const [syncTarget, setsyncTarget] = useState<string>("");

  const [data, setdata] = useState<Request[]>([]);

  const syncRequestRecord = (id: string) => {
    setmodalTitle("Sync Request Record");
    setmodalBody("คุณต้องการที่จะ Sync ข้อมูลนี้ใช่หรือไม่?");
    setsyncTarget(id);
    setmodalConfirm(() => () => confirmSyncRequestRecord());
    onOpen();
  };

  const confirmSyncRequestRecord = async () => {
    await syncRequestRecordService(syncTarget)
      .then(() => {
        addToast({
          description: "Sync ข้อมูลสำเร็จ",
          status: ToastStatus.SUCCESS,
        });
      })
      .catch(() => {
        addToast({
          description: "เกิดข้อผิดพลาดขณะทำการ Sync ข้อมูล",
          status: ToastStatus.ERROR,
        });
      });
    onClose();
  };

  const syncAllRequestRecord = () => {
    setmodalTitle("Sync All Request Record");
    setmodalBody(
      "คุณต้องการที่จะ Sync ข้อมูลทั้งหมดใช่หรือไม่? ระบบจะเลือกเฉพาะคำขอล่าสุดของแต่ละ record ที่อยู่ในสถานะ `ตรวจสอบเรียบร้อยแล้ว` เท่านั้น"
    );
    setmodalConfirm(() => () => confirmSyncAllRequestRecord());
    onOpen();
  };

  const confirmSyncAllRequestRecord = async () => {
    await syncAllRequestRecordService()
      .then(() => {
        addToast({
          description: "Sync ข้อมูลสำเร็จ",
          status: ToastStatus.SUCCESS,
        });
      })
      .catch(() => {
        addToast({
          description: "เกิดข้อผิดพลาดขณะทำการ Sync ข้อมูล",
          status: ToastStatus.ERROR,
        });
      });

    onClose();
  };

  useEffect(() => {
    (async () => {
      const isAuthorize = await AuthorizeAdmin(Role.ADMIN);
      if (isAuthorize === false) {
        navigate("/user/login");
      }
    })();

    const getRequest = async () => {
      await getRequestByParamsService({})
        .then((res) => {
          const r = res.map((item) => {
            item.startTime = extractStringFromHTML(item.startTime);
            item.endTime = extractStringFromHTML(item.endTime);
            item.question = extractStringFromHTML(item.question);
            item.answer = extractStringFromHTML(item.answer);
            item.createdAt = new Date(item.createdAt).toLocaleString("th-TH");
            item.status = convertStatusWord(item.status);
            return item;
          });

          setdata(r);
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
    getRequest();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <AdminBasePage activePage="Request">
      <Flex w="full" direction="column" gap={4}>
        <Flex alignSelf={"flex-end"} p={4}>
          <Button colorScheme="teal" onClick={syncAllRequestRecord}>
            Sync All
          </Button>
        </Flex>
        <RequestTable data={data} syncRequestRecord={syncRequestRecord} />
      </Flex>

      <ActionModal
        openModal={isOpen}
        closeModal={onClose}
        modalTitle={modalTitle}
        modalBody={modalBody}
        confirm={modalConfirm}
      />
    </AdminBasePage>
  );
}

export default AdminRequestPage;

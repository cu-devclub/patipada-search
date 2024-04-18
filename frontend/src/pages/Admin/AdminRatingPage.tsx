import { useEffect, useState } from "react";
import { AdminBasePage } from "./AdminBasePage";
import { useNavigate } from "react-router-dom";
import { MessageToast } from "../../components";
import { Role, ToastStatus } from "../../constant";
import { AuthorizeAdmin } from "../../functions";
import { Flex } from "@chakra-ui/react";
import { FullRating } from "../../models/ratings";
import { getRatingsService } from "../../service/data/ratings";
import { RatingTable } from "../../components/rating";

function AdminRatingPage() {
  const navigate = useNavigate();
  const { addToast } = MessageToast();

  const [ratings, setRatings] = useState<FullRating[]>([]);

  useEffect(() => {
    (async () => {
      const isAuthorize = await AuthorizeAdmin(Role.ADMIN);
      if (isAuthorize === false) {
        navigate("/user/login");
      }
    })();
    const getRatings = async () => {
      await getRatingsService()
        .then((res: FullRating[]) => {
          setRatings(res);
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
    getRatings();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <AdminBasePage activePage="Ratings">
      <Flex w="full" direction="column" gap={4}>
        <RatingTable ratings={ratings} />
      </Flex>
    </AdminBasePage>
  );
}

export default AdminRatingPage;

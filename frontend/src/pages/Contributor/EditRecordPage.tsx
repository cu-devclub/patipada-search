import { useParams } from "react-router-dom";
import { search } from "../../service/search";
import { useEffect,useState } from "react";
import {
  EditRecordHeader,
  EditRecordForm,
} from "../../components/contributor/edit-record";
import { Grid } from "@chakra-ui/react";
import { Footer } from "../../components";
import { DataItem } from "../../models/qa";
function EditRecordPage() {
  const { recordID } = useParams();
  const [data, setData] = useState<DataItem>()

  useEffect(() => {
    const getRecord = async (recordID: string) => {
      await search(recordID)
        .then((res) => {
          setData(res.data[0]);
        })
        .catch((err) => {
          console.log(err);
        });
    };
    if (recordID) getRecord(recordID);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [recordID]);

  return (
    <Grid templateRows="auto 1fr auto" gap={4} w="full" minH="100svh">
      <EditRecordHeader />
      {data && <EditRecordForm data={data}/>}
      <Footer />
    </Grid>
  );
}

export default EditRecordPage;

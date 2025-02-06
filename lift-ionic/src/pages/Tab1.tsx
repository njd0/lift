import {
  IonContent,
  IonHeader,
  IonPage,
  IonTitle,
  IonToolbar,
} from "@ionic/react"
import ExploreContainer from "../components/ExploreContainer"
import "./Tab1.css"
import { useGetAllExercises } from "../api/ninjas/exercises/hooks"
import Calendar from "../components/Calendar"

const Tab1: React.FC = () => {
  const { data, isLoading } = useGetAllExercises({
    params: {
      name: "barbell",
    },
  })

  return (
    <IonPage>
      <IonHeader>
        <IonToolbar>
          <IonTitle>Tab 1</IonTitle>
        </IonToolbar>
      </IonHeader>
      <IonContent fullscreen>
        <IonHeader collapse='condense'>
          <IonToolbar>
            <IonTitle size='large'>Tab 1</IonTitle>
          </IonToolbar>
        </IonHeader>
        {/* <ExploreContainer name='Tab 1 page' /> */}
        {/* <Calendar /> */}
      </IonContent>
    </IonPage>
  )
}

export default Tab1

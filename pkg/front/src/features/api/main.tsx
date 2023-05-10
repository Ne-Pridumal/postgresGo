import { ROOT_URL } from "@shared/consts"
import axios from "axios"
import { useMutation, useQuery } from "react-query"


const QUERY_KEY = "get-all-bookings"

type TQuery = {
  amount: number,
  name: string,
  departureDate: Date,
}
type TParams = {
  destination: string,
  departure: string,
  date: string,
  price: number,
  name: string,
}

const fetcher = async (props: TParams): Promise<TQuery> => {
  const date = Date.now()
  const bookingData = await fetch(ROOT_URL + `/bookings`, {
    body: JSON.stringify({
      ref: date,
      date: new Date(props.date).toISOString(),
      amount: Number(props.price),
    }),
    method: 'POST',
  })
  const ticketData = await fetch(ROOT_URL + `/tickets`, {
    body: JSON.stringify(
      {
        key: date.toString(),
        bookRef: date,
        passengerId: date,
        passengerName: props.name,
        contactDate: new Date().toISOString(),
      }
    ),
    method: "POST"
  })
  const flightData = await fetch(ROOT_URL + `/flights`, {
    body: JSON.stringify({
      key: date,
      scheduledDeparture: new Date(props.date).toISOString(),
      scheduledArrival: new Date(props.date).toISOString(),
      departureAirport: 25459237504293,
      arrivalAirport: 25459237504293,
      status: "prep",
      aircraftKey: 3452453,
      actualDeparture: new Date(props.date).toISOString(),
      actualArrival: new Date(props.date).toISOString(),
    }),
    method: "POST"
  })
  const ticketFlightData = await fetch(ROOT_URL + `/ticketFlights`, {
    body: JSON.stringify({
      flightKey: date,
      scheduledDeparture: new Date(props.date).toISOString(),
      ticketKey: date.toString(),
      fareCondition: "fdsfsdfs",
      amount: Number(props.price),
    }),
    method: "POST"
  })
  const boardingPassData = await fetch(ROOT_URL + `/boardingPasses`, {
    body: JSON.stringify({
      flightKey: date,
      scheduledDeparture: new Date(props.date).toISOString(),
      ticketKey: date.toString(),
      boardingKey: date.toString(),
      seatKey: "13A"
    }),
    method: "POST"
  })
  console.log({ bookingData, ticketData, ticketFlightData, boardingPassData })
  const response: TQuery = {
    name: props.name,
    amount: props.price,
    departureDate: new Date(date)
  }
  return response
}

export const useCreateTicket = () => useMutation(
  QUERY_KEY,
  (p: TParams) => fetcher(p),
  {}
)

export const useTestQuery = () => useQuery(
  'asdfasd',
  testFetch,
  {}
)

const testFetch = async () => {
  const { data } = await axios.get(ROOT_URL + `/bookings`)
  return data
}

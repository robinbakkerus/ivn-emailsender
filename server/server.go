package server

func Init() {
  r := router.Router()
  r.Run(":3000")
}
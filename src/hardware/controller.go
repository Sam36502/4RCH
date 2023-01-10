package hardware

type Controller struct {
	KeyUp    int32 `json:"kc_up"`
	KeyDown  int32 `json:"kc_down"`
	KeyLeft  int32 `json:"kc_left"`
	KeyRight int32 `json:"kc_right"`
	KeyA     int32 `json:"kc_a"`
	KeyB     int32 `json:"kc_b"`
	KeyC     int32 `json:"kc_c"`
	KeyD     int32 `json:"kc_d"`
}

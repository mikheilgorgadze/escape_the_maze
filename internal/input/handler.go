package input

import (
	"github.com/eiannone/keyboard"
)

type Movement struct{
    Row int
    Col int
}

type Handler struct{
    KeyMappings map[keyboard.Key]Movement
    KeyEvents   <- chan keyboard.KeyEvent
}

func NewHandler() (*Handler, error){
    keyEvents, err := keyboard.GetKeys(100)
    if err!=nil{
        return nil, err
    }

    return &Handler{
        KeyMappings: map[keyboard.Key]Movement{
            keyboard.KeyArrowRight: {0, 1},
            keyboard.KeyArrowLeft: {0, -1},
            keyboard.KeyArrowDown: {1, 0},
            keyboard.KeyArrowUp: {-1, 0},
        },
        KeyEvents: keyEvents,
    }, nil
}

func (h *Handler) Close(){
    keyboard.Close()
}

func (h *Handler) GetMovement(key keyboard.Key) (Movement, bool){
    movement, exists := h.KeyMappings[key]
    return movement, exists
}

func (h *Handler) IsExit(key keyboard.Key) bool{
    return key == keyboard.KeyEsc
}

func (h *Handler) IsStart(key keyboard.Key) bool{
    return key == keyboard.KeyEnter
}

func (h *Handler) ReadKey() (keyboard.KeyEvent, error){
    events := <- h.KeyEvents
    return events, events.Err
}


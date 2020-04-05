/**
* @file event_dispatcher.go
* @brief A simple, machine-local, thread-safe pub/sub event queue and dispatcher for golang.
* @author Anadian
* @copyright 	Copyright 2019 Canosw
	Permission is hereby granted, free of charge, to any person obtaining a copy of this 
software and associated documentation files (the "Software"), to deal in the Software 
without restriction, including without limitation the rights to use, copy, modify, 
merge, publish, distribute, sublicense, and/or sell copies of the Software, and to 
permit persons to whom the Software is furnished to do so, subject to the following 
conditions:
	The above copyright notice and this permission notice shall be included in all copies 
or substantial portions of the Software.
	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, 
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A 
PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT 
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF 
CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE 
OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package event_dispatcher;

//# Dependencies
import(
	//## Internal
	//## Standard
	"time"
	"sync"
	"strconv"
	//## External
	error_report "github.com/Anadian/error_report/source"
	matchkey "github.com/Anadian/matchkey/source"
);

//# Constants
const(
	//## Exported Constants
	//### Errors
	ERROR_CODE_MATCH_ERROR int64 = 18;
	ERROR_CODE_INDEX_OUT_OF_RANGE int64 = 20;
	ERROR_CODE_SUBORDINATE_FUNCTION_ERROR int64 = 4;
	ERROR_CODE_INVALID_MATCHKEY_TYPE int64 = 5;
	ERROR_CODE_EVENT_LISTENER_MATCH int64 = 6;
	ERROR_CODE_EVENT_PROCESSING_ERROR int64 = 7;
	//## Private Constants
);

//Types, structs, and methods
type Event_struct struct{
	name string
	//time time.Time
	data map[string]interface{}
}
type EventListener_struct struct{
	key matchkey.MatchKey_struct
	async bool
	function func( event Event_struct, args ...interface{} )
}
type EventDispatcher_struct struct{
	mutex sync.Mutex
	add_times bool
	buffered bool
	events_slice []Event_struct
	event_listeners_slice []EventListener_struct
}

/**
* @fn AddEventListener
* @brief Adds an event listener to the event dispatcher.
* @struct event_dispatcher *EventDispatcher_struct
* @param event_listener EventListener_struct [in] The event listener object to add.
* @return (return_report error_report.ErrorReport_Struct)
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// AddEventListener adds an event listener to the event dispatcher.
func (event_dispatcher *EventDispatcher_struct) AddEventListener( event_listener EventListener_struct ) (return_report error_report.ErrorReport_struct){
	/* Variables */
	/* Parametres */
	/* Function */
	event_dispatcher.mutex.Lock();
	event_dispatcher.event_listeners_slice = append(event_dispatcher.event_listeners_slice, event_listener);
	return_report = error_report.New( 0, map[string]interface{}{ "event_listeners_slice_length": len(event_dispatcher.event_listeners_slice) }, nil );
	event_dispatcher.mutex.Unlock();
	/* Return */
	return return_report;
}

/**
* @fn RemoveEventListenerByStringLiteral
* @brief RemoveEventListenerByStringLiteralByStringLiteral
* @struct event_dispatcher *EventDispatcher_struct
* @param string_literal string [in] The key.matchkey_string value for the event listener to be removed.
* @return (return_report error_report.ErrorReport_struct)
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// RemoveEventListenerByStringLiteral removeEventListenerByStringLiteral
func (event_dispatcher *EventDispatcher_struct) RemoveEventListenerByStringLiteral( string_literal string ) (return_report error_report.ErrorReport_struct){
	/* Variables */
	var i uint;
	var before_slice []EventListener_struct;
	var after_slice []EventListener_struct;
	/* Parametres */
	/* Function */
	event_dispatcher.mutex.Lock();
	for i = 0; int(i) < len(event_dispatcher.event_listeners_slice); i++ {
		if( event_dispatcher.event_listeners_slice[i].key.Matchkey_string == string_literal ){
			before_slice = event_dispatcher.event_listeners_slice[:i];
			after_slice = event_dispatcher.event_listeners_slice[(i+1):];
			event_dispatcher.event_listeners_slice = append(before_slice,after_slice...);
		}
	}
	return_report = error_report.New( 0, map[string]interface{}{ "event_listeners_slice_length": len(event_dispatcher.event_listeners_slice) }, nil );
	event_dispatcher.mutex.Unlock();
	/* Return */
	return return_report;
}

/**
* @fn GetEventByIndex
* @brief Returns a copy of the event, at the given index, in the events slice; it does not modify the event slice.
* @struct event_dispatcher EventDispatcher_struct
* @param index uint [in] The index you the event to be retrieved.
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// GetEventByIndex returns a copy of the event, at the given index, in the events slice; it does not modify the event slice.
func (event_dispatcher EventDispatcher_struct) GetEventByIndex( index uint ) ( return_report error_report.ErrorReport_struct ){
	//Variables
	var event Event_struct;
	//Parametres
	//Function
	event_dispatcher.mutex.Lock();
	if( int(index) < len(event_dispatcher.events_slice) ){
		event = event_dispatcher.events_slice[index];
		return_report = error_report.New( 0, map[string]interface{}{ "event": event }, nil );
	} else{
		return_report = error_report.New( ERROR_CODE_INDEX_OUT_OF_RANGE, map[string]interface{}{ "message": "index out of range.", "events_slice_length": len(event_dispatcher.events_slice) }, nil );
	}
	event_dispatcher.mutex.Unlock();

	//Return
	return return_report;
}

/**
* @fn RemoveEventByIndex
* @brief Removes the event at the given index from the events slice.
* @struct event_dispatcher *EventDispatcher_struct
* @param index uint [in] The index of the event to remove.
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// RemoveEventByIndex removes the event at the given index from the events slice.
func (event_dispatcher *EventDispatcher_struct) RemoveEventByIndex( index uint ) ( return_report error_report.ErrorReport_struct ){
	//Variables
	var preceding_events []Event_struct;
	var following_events []Event_struct;
	//Parametres
	//Function
	event_dispatcher.mutex.Lock();
	if( int(index) < len(event_dispatcher.events_slice) ){
		preceding_events = event_dispatcher.events_slice[:index];
		following_events = event_dispatcher.events_slice[(index+1):];
		event_dispatcher.events_slice = append(preceding_events, following_events...);
		return_report = error_report.New( 0, map[string]interface{}{ "events_slice_length": len(event_dispatcher.events_slice) }, nil );
	} else{
		return_report = error_report.New( ERROR_CODE_INDEX_OUT_OF_RANGE, map[string]interface{}{ "message": "index out of range.", "events_slice_length": len(event_dispatcher.events_slice) }, nil );
	}
	event_dispatcher.mutex.Unlock();
	//Return
	return return_report;
}

/**
* @fn ExtractEventByIndex
* @brief Extracts the event at the given index from the events slice, removing it from the slice, and returns it.
* @struct event_dispatcher *EventDispatcher_struct
* @param index uint [in] The index of the event to be extracted.
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// ExtractEventByIndex extracts the event at the given index from the events slice, removing it from the slice, and returns it.
func (event_dispatcher *EventDispatcher_struct) ExtractEventByIndex( index uint ) ( return_report error_report.ErrorReport_struct ){
	//Variables
	var temp_event interface{};
	var event Event_struct;
	var get_error_report error_report.ErrorReport_struct;
	var remove_error_report error_report.ErrorReport_struct;
	//Parametres
	//Function
	get_error_report = event_dispatcher.GetEventByIndex(index);
	if( get_error_report.NoError() == true ){
		temp_event = get_error_report.Data["event"]
		event = temp_event.(Event_struct);
		remove_error_report = event_dispatcher.RemoveEventByIndex(index);
		if( remove_error_report.NoError() == true ){
			return_report = error_report.New( 0, map[string]interface{}{ "event": event, "events_slice_length": remove_error_report.Data["events_slice_length"] }, nil );
		} else{
			return_report = error_report.New( ERROR_CODE_SUBORDINATE_FUNCTION_ERROR, map[string]interface{}{ "message": "event_dispatcher.RemoveEventByIndex() returned an error." }, &remove_error_report );
		}
	} else{
		return_report = error_report.New( ERROR_CODE_SUBORDINATE_FUNCTION_ERROR, map[string]interface{}{ "message": "event_dispatcher.GetEventByIndex() returned an error." }, &get_error_report );
	}
	//Return
	return return_report;
}

/**
* @fn InsertEventAtIndex
* @brief Inserts the given event into the events slice at the given index.
* @struct event_dispatcher *EventDispatcher_struct
* @param event Event_struct [in] 
* @param index uint [in] 
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// InsertEventAtIndex inserts the given event into the events slice at the given index.
func (event_dispatcher *EventDispatcher_struct) InsertEventAtIndex( event Event_struct, index uint ) ( return_report error_report.ErrorReport_struct ){
	//Variables
	var preceding_events []Event_struct;
	var following_events []Event_struct;
	//Parametres
	//Function
	event_dispatcher.mutex.Lock();
	if( event_dispatcher.add_times == true ){
		event.data["submission_time"] = time.Now();
	}
	if( int(index) >= len(event_dispatcher.events_slice) ){
		event_dispatcher.events_slice = append(event_dispatcher.events_slice,event);
	} else{
		preceding_events = event_dispatcher.events_slice[:index];
		following_events = event_dispatcher.events_slice[(index+1):];
		event_dispatcher.events_slice = append(preceding_events,event);
		event_dispatcher.events_slice = append(event_dispatcher.events_slice,following_events...);
	}
	return_report = error_report.New( 0, map[string]interface{}{ "events_slice_length": len(event_dispatcher.events_slice) }, nil );
	event_dispatcher.mutex.Unlock();
	//Return
	return return_report;
}

/**
* @fn PublishEvent
* @brief Publishes the given event to listeners registered with the dispatcher.
* @struct event_dispatcher EventDispatcher_struct
* @param event Event_struct [in] The event to be published.
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// PublishEvent publishes the given event to listeners registered with the dispatcher.
func (event_dispatcher EventDispatcher_struct) PublishEvent( event Event_struct ) ( return_report error_report.ErrorReport_struct ){
	//Variables
	var listener_index int;
	var match bool;
	var match_error_report error_report.ErrorReport_struct;
	var listener_index_string string;
	var errors_interface interface{};
	var errors_int int;
	//Parametres
	//Function
	event_dispatcher.mutex.Lock();
	if( event_dispatcher.buffered == true ){
		event_dispatcher.events_slice = append(event_dispatcher.events_slice, event);
	} else{
		for listener_index = 0; listener_index < len(event_dispatcher.event_listeners_slice); listener_index++ {
			match, match_error_report = event_dispatcher.event_listeners_slice[listener_index].key.Match( event.name );
			if( match_error_report.NoError() == true ){
				if( match == true ){
					if( event_dispatcher.event_listeners_slice[listener_index].async == true ){
						go event_dispatcher.event_listeners_slice[listener_index].function( event );
					} else{
						event_dispatcher.event_listeners_slice[listener_index].function( event );
					}
				}
			} else{
				listener_index_string = strconv.FormatUint( uint64(listener_index), 10 );
				return_report.Code = ERROR_CODE_MATCH_ERROR;
				errors_interface = return_report.Data["errors"];
				errors_int = errors_interface.(int);
				errors_int++;
				return_report.Data["errors"] = errors_int;
				return_report.Data[listener_index_string] = match_error_report;
			}
		}
	}
	event_dispatcher.mutex.Unlock();
	//Return
	return return_report;
}

/**
* @fn PushEvent
* @brief Adds an event to the end of the event queue.
* @struct event_dispatcher *EventDispatcher_struct
* @param event Event_struct [in] The event to be added to the end of the queue.
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// PushEvent adds an event to the end of the event queue.
func (event_dispatcher *EventDispatcher_struct) PushEvent( event Event_struct ) ( return_report error_report.ErrorReport_struct ){
	//Variables
	//Parametres
	//Function
	event_dispatcher.mutex.Lock();
	if( event_dispatcher.add_times == true ){
		event.data["submission_time"] = time.Now();
	}
	event_dispatcher.events_slice = append(event_dispatcher.events_slice, event);
	return_report = error_report.New( 0, map[string]interface{}{ "new_length": len(event_dispatcher.events_slice) }, nil );
	event_dispatcher.mutex.Unlock();
	//Return
	return return_report;
}

/**
* @fn PopEvent
* @brief Returns the last event in the queue.
* @struct event_dispatcher *EventDispatcher
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// PopEvent returns the last event in the queue.
func (event_dispatcher *EventDispatcher_struct) PopEvent() ( return_report error_report.ErrorReport_struct ){
	//Variables
	var event Event_struct;
	//Parametres
	//Function
	event_dispatcher.mutex.Lock();
	event = event_dispatcher.events_slice[(len(event_dispatcher.events_slice) - 1)];
	event_dispatcher.events_slice = event_dispatcher.events_slice[:(len(event_dispatcher.events_slice) - 1)];
	event_dispatcher.mutex.Unlock();
	return_report = error_report.New( 0, map[string]interface{}{ "event": event }, nil );
	//Return
	return return_report;
}

/**
* @fn ShiftEvent
* @brief Extracts and returns the first event in the queue.
* @struct event_dispatcher EventDispatcher_struct
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// ShiftEvent extracts and returns the first event in the queue.
func (event_dispatcher EventDispatcher_struct) ShiftEvent() ( return_report error_report.ErrorReport_struct ){
	//Variables
	var function_return error_report.ErrorReport_struct;
	//Parametres
	//Function
	function_return = event_dispatcher.ExtractEventByIndex( 0 );
	if( function_return.NoError() == true ){
		return_report = error_report.New( 0, map[string]interface{}{ "event": function_return.Data["event"] }, &function_return );
	} else{
		return_report = error_report.New( ERROR_CODE_SUBORDINATE_FUNCTION_ERROR, map[string]interface{}{}, &function_return );
	}
	//Return
	return return_report;
}

/**
* @fn ProcessEvent
* @brief Transmits the given event.
* @struct event_dispatcher EventDispatcher_struct
* @param event Event_struct [in] The event to be processed.
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// ProcessEvent transmits the given event.
func (event_dispatcher EventDispatcher_struct) ProcessEvent( event Event_struct ) ( return_report error_report.ErrorReport_struct ){
	//Variables
	var function_return error_report.ErrorReport_struct;
	//Parametres
	event_dispatcher.mutex.Lock();
	function_return = event_dispatcher.ProcessEvent_Unsafe( event );
	if( function_return.IsError() == true ){
		return_report = error_report.New( ERROR_CODE_EVENT_PROCESSING_ERROR, map[string]interface{}{ "event": event }, &function_return );
	}
	event_dispatcher.mutex.Unlock();
	//Return
	return return_report;
}

/**
* @fn ProcessEvent_Unsafe
* @brief Actually transmits the event.
* @struct event_dispatcher EventDispatcher_struct
* @param event Event_struct [in] The event to be transmitted.
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// ProcessEvent_Unsafe actually transmits the event.
func (event_dispatcher EventDispatcher_struct) ProcessEvent_Unsafe( event Event_struct ) ( return_report error_report.ErrorReport_struct ){
	//Variables
	var i int; //Event listener index
	var match bool;
	var index_string string;
	var function_return error_report.ErrorReport_struct;
	//Parametres
	if( event_dispatcher.add_times == true ){
		event.data["transmission_time"] = time.Now();
	}
	//Function
	for i = 0; i < len(event_dispatcher.event_listeners_slice); i++ {
		match, function_return = event_dispatcher.event_listeners_slice[i].key.Match( event.name );
		if( function_return.NoError() == true ){
			if( match == true ){
				if( event_dispatcher.event_listeners_slice[i].async == true ){
					go event_dispatcher.event_listeners_slice[i].function( event );
				} else{
					event_dispatcher.event_listeners_slice[i].function( event );
				}
			}
		} else{
			index_string = strconv.FormatUint( uint64(i), 10 );
			if( return_report.CodeEqual( 0 ) == true ){
				return_report = error_report.New( ERROR_CODE_EVENT_LISTENER_MATCH, map[string]interface{}{ "message": "Matching against an event listener returned an error.", index_string: function_return }, nil);
			} else{
				return_report.Data[index_string] = function_return;
			}
		}
	}
	//Return
	return return_report;
}

/**
* @fn ProcessEvents
* @brief Processes all of the events in the queue.
* @struct event_dispatcher EventDispatcher_struct
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// ProcessEvents processes all of the events in the queue.
func (event_dispatcher EventDispatcher_struct) ProcessEvents() ( return_report error_report.ErrorReport_struct ){
	//Variables
	var function_return error_report.ErrorReport_struct;
	var event Event_struct;
	//Parametres
	//Function
	function_return = event_dispatcher.ShiftEvent();
	if( function_return.NoError() == true ){
		event = function_return.Data["event"].(Event_struct);
		function_return = event_dispatcher.ProcessEvent( event );
		if( function_return.NoError() == true ){
			return_report = error_report.New( 0, map[string]interface{}{ "event": event }, nil );
		} else{
			return_report = error_report.New( ERROR_CODE_EVENT_PROCESSING_ERROR, map[string]interface{}{}, &function_return );
		}
	} else{
		return_report = error_report.New( ERROR_CODE_SUBORDINATE_FUNCTION_ERROR, map[string]interface{}{}, &function_return );
	}
	//Return
	return return_report;
}


/*type EventEmitter_struct struct{
	events_slice []Event_struct
	listeners_map map[string]func(args ...interface{})
	mutex sync.Mutex
}*/

/**
* @fn AddListener
* @brief Adds a listener function to the event emitter.
* @struct event_emitter *EventEmitter_struct
* @param event_string string [in] A string identifying which events will trigger the `listener_function` to be called.
* @param listener_function func(args ...interface{}) [in] The function to call when an event matches `event_string`.
* @param default_arguments ...interface{} [in] The default parametre arguments used when the `listener_function` is triggered but the event doesn't specify any additional arguments.
* @return error
* @retval nil Success
* @retval 1 Not Supported
* @retval !== nil && !== 1 Error
*/
// AddListener adds a listener function to the event emitter.
/*func (event_emitter *EventEmitter_struct) AddListener( event_string string, listener_function func(args ...interface{}), default_arguments ...interface{} ) error{
	var _return error = nil;
	/* Variables */
	/* Parametres */
	/* Function */
	/* Return */
	/*return _return;
}*/



//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//# Exported Functions
/**
* @fn NewEvent
* @brief Creates a new event.
* @param name string [in] The name of the event.
* @param data map[string]interface{} [in] A string-keyed map of extra data contained in the event.
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// NewEvent creates a new event.
func NewEvent( name string, data map[string]interface{} ) ( return_report error_report.ErrorReport_struct ){
	//Variables
	var event Event_struct;
	//Parametres
	//Function
	event.name = name;
	event.data = data;
	event.data["creation_time"] = time.Now();
	return_report = error_report.New( 0, map[string]interface{}{ "event": event }, nil );
	//Return
	return return_report;
}

/**
* @fn NewEventListener
* @brief Creates a new event listener.
* @param key matchkey.Matchkey_struct [in] The Matchkey_struct to trigger the event listener.
* @param async bool [in] A boolean expressing whether the event listner function should be called in its own go routine.
* @param function func( event Event_struct, args ...interface{}) [in] The function to be called when the event matches the matchkey.
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// NewEventListener creates a new event listener.
func NewEventListener( key matchkey.MatchKey_struct, async bool, function func( event Event_struct, args ...interface{}) ) ( return_report error_report.ErrorReport_struct ){
	//Variables
	var event_listener EventListener_struct;
	//Parametres
	//Function
	if( (key.Matchkey_type > 0) && (key.Matchkey_type <= 3) ){
		event_listener.key = key;
		event_listener.async = async;
		event_listener.function = function;
		return_report = error_report.New( 0, map[string]interface{}{ "event_listener": event_listener }, nil );
	} else{
		return_report = error_report.New( ERROR_CODE_INVALID_MATCHKEY_TYPE, map[string]interface{}{ "message": "Invalid Matchkey_type property for the given Matchkey_struct `key` argument." }, nil );
	}
	//Return
	return return_report;
}

/**
* @fn NewEventDispatcher
* @brief Creates a new event dispatcher.
* @param add_times bool [in] A boolean value representing whether to add submission and transmission times to events.
* @param buffered bool [in] A boolean value representing whther to queue events and only actually transmit them when `ProcessEvents` is called manually.
* @return ( return_report error_report.ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// NewEventDispatcher creates a new event dispatcher.
func NewEventDispatcher( add_times bool, buffered bool ) ( return_report error_report.ErrorReport_struct ){
	//Variables
	var event_dispatcher EventDispatcher_struct;
	//Parametres
	//Function
	event_dispatcher.add_times = add_times;
	event_dispatcher.buffered = buffered;
	return_report = error_report.New( 0, map[string]interface{}{ "event_dispatcher": event_dispatcher }, nil );
	//Return
	return return_report;
}

//Private Functions


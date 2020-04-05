/**
* @file event_dispatcher_test.go
* @brief Contains test functions for `event_dispatcher.go`.
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

// event_dispatcher_test contains test functions for `event_dispatcher.go`.
package event_dispatcher;

//# Dependencies
import(
	//## Internal
	//## Standard
	"testing"
	"log"
	//## External
	error_report "github.com/Anadian/error_report/source"
	matchkey "github.com/Anadian/matchkey/source"
);

//# Constants
const(
	//## Exported Constants
	//### Errors
	//## Private Constants
);

//# Types
//## Interfaces
//## Structs
//### Methods

//# Global Variables
var(
	//## Exported Variables
	//## Private Variables
);

//# Exported Functions
/**
* @fn TestEventDispatcher
* @brief Tests EventDispatcher_struct's methods.
* @param t *testing.T [in] Go stdlib testing object.
*/

// TestEventDispatcher tests EventDispatcher_struct's methods.
func TestEventDispatcher( t *testing.T ){
	//Variables
	var event_dispatcher EventDispatcher_struct;
	var string_event_listener_matchkey matchkey.MatchKey_struct;
	var path_event_listener_matchkey matchkey.MatchKey_struct;
	var regex_event_listener_matchkey matchkey.MatchKey_struct;
	var event_listener EventListener_struct;
	var event Event_struct;
	var function_return, function_return2, function_return3 error_report.ErrorReport_struct;

	//Parametres
	//Function
	log.SetFlags( log.LstdFlags | log.Lmicroseconds | log.Lshortfile );
	///Creating matchkeys
	string_event_listener_matchkey, function_return = matchkey.New( matchkey.MATCHKEY_TYPE_STRING, "event_listener:string_test" );
	path_event_listener_matchkey, function_return2 = matchkey.New( matchkey.MATCHKEY_TYPE_PATH, "event_listener:*_test" );
	regex_event_listener_matchkey, function_return3 = matchkey.New( matchkey.MATCHKEY_TYPE_REGEX, "event_listener:[a-zA-Z]*[_-]test" );
	if( function_return.NoError() && function_return2.NoError() && function_return3.NoError() ){
		log.Printf("Success: No errors when creating needed matchkeys.\n");
	} else{
		t.Fail();
		log.Printf("Failure: matchkey.New returned an error: string: %v, path: %v, regex: %v\n", function_return, function_return2, function_return3);
	}
	//Creating event dispatcher (true, true)
	function_return = NewEventDispatcher( true, true );
	if( function_return.NoError() == true ){
		log.Printf("Success: NewEventDispatcher returned no errors.\n");
		event_dispatcher = function_return.Data["event_dispatcher"].(EventDispatcher_struct);
		//Create event listener string (true)
		function_return = NewEventListener( string_event_listener_matchkey, true, func( event Event_struct, args ...interface{} ){
			log.Printf("event_listener_string (true) received event: %v\n", event);
		} );
		if( function_return.NoError() == true ){
			log.Printf("Success: Event listener (string true) created.\n");
			event_listener = function_return.Data["event_listener"].(EventListener_struct);
			//Add event listener 
			function_return = event_dispatcher.AddEventListener( event_listener );
			if( function_return.NoError() == true ){
				log.Printf("Success: string (true) event listener added.\n");
			} else{
				t.Fail();
				log.Printf("Failure: event_dispatcher.AddEventListener returned an error: %v\n", function_return);
			}
		} else{
			t.Fail();
			log.Printf("Failure: NewEventListener returned an error: %v\n", function_return);
		}
		//Create event listener path (true)
		function_return = NewEventListener( path_event_listener_matchkey, true, func( event Event_struct, args ...interface{} ){
			log.Printf("event_listener_path (true) received event: %v\n", event);
		} );
		if( function_return.NoError() == true ){
			log.Printf("Success: Event listener (path true) created.\n");
			event_listener = function_return.Data["event_listener"].(EventListener_struct);
			//Add event listener 
			function_return = event_dispatcher.AddEventListener( event_listener );
			if( function_return.NoError() == true ){
				log.Printf("Success: path (true) event listener added.\n");
			} else{
				t.Fail();
				log.Printf("Failure: event_dispatcher.AddEventListener returned an error: %v\n", function_return);
			}
		} else{
			t.Fail();
			log.Printf("Failure: NewEventListener returned an error: %v\n", function_return);
		}
		//Create event listener regex (true)
		function_return = NewEventListener( regex_event_listener_matchkey, true, func( event Event_struct, args ...interface{} ){
			log.Printf("event_listener_regex (true) received event: %v\n", event);
		} );
		if( function_return.NoError() == true ){
			log.Printf("Success: Event listener (regex true) created.\n");
			event_listener = function_return.Data["event_listener"].(EventListener_struct);
			//Add event listener 
			function_return = event_dispatcher.AddEventListener( event_listener );
			if( function_return.NoError() == true ){
				log.Printf("Success: regex (true) event listener added.\n");
			} else{
				t.Fail();
				log.Printf("Failure: event_dispatcher.AddEventListener returned an error: %v\n", function_return);
			}
		} else{
			t.Fail();
			log.Printf("Failure: NewEventListener returned an error: %v\n", function_return);
		}
		//Create event listener string (false)
		function_return = NewEventListener( string_event_listener_matchkey, false, func( event Event_struct, args ...interface{} ){
			log.Printf("event_listener_string (false) received event: %v\n", event);
		} );
		if( function_return.NoError() == true ){
			log.Printf("Success: Event listener (string false) created.\n");
			event_listener = function_return.Data["event_listener"].(EventListener_struct);
			//Add event listener 
			function_return = event_dispatcher.AddEventListener( event_listener );
			if( function_return.NoError() == true ){
				log.Printf("Success: string (false) event listener added.\n");
			} else{
				t.Fail();
				log.Printf("Failure: event_dispatcher.AddEventListener returned an error: %v\n", function_return);
			}
		} else{
			t.Fail();
			log.Printf("Failure: NewEventListener returned an error: %v\n", function_return);
		}
		//Create event listener path (false)
		function_return = NewEventListener( path_event_listener_matchkey, false, func( event Event_struct, args ...interface{} ){
			log.Printf("event_listener_path (false) received event: %v\n", event);
		} );
		if( function_return.NoError() == true ){
			log.Printf("Success: Event listener (path false) created.\n");
			event_listener = function_return.Data["event_listener"].(EventListener_struct);
			//Add event listener 
			function_return = event_dispatcher.AddEventListener( event_listener );
			if( function_return.NoError() == true ){
				log.Printf("Success: path (false) event listener added.\n");
			} else{
				t.Fail();
				log.Printf("Failure: event_dispatcher.AddEventListener returned an error: %v\n", function_return);
			}
		} else{
			t.Fail();
			log.Printf("Failure: NewEventListener returned an error: %v\n", function_return);
		}
		//Create event listener regex (false)
		function_return = NewEventListener( regex_event_listener_matchkey, false, func( event Event_struct, args ...interface{} ){
			log.Printf("event_listener_regex (false) received event: %v\n", event);
		} );
		if( function_return.NoError() == true ){
			log.Printf("Success: Event listener (regex false) created.\n");
			event_listener = function_return.Data["event_listener"].(EventListener_struct);
			//Add event listener 
			function_return = event_dispatcher.AddEventListener( event_listener );
			if( function_return.NoError() == true ){
				log.Printf("Success: regex (false) event listener added.\n");
			} else{
				t.Fail();
				log.Printf("Failure: event_dispatcher.AddEventListener returned an error: %v\n", function_return);
			}
		} else{
			t.Fail();
			log.Printf("Failure: NewEventListener returned an error: %v\n", function_return);
		}
		//Create event string_test
		function_return = NewEvent( "event_listener:string_test", map[string]interface{}{} );
		if( function_return.NoError() == true ){
			log.Printf("Success: Event event_lister:string_test created successfully.\n",);
			event = function_return.Data["event"].(Event_struct);
			function_return = event_dispatcher.PushEvent(event);
			if( function_return.NoError() == true ){
				log.Printf("Success: Event string_test pushed successfully.\n",);
			} else{
				log.Printf("Failure: event_dispatcher.PushEvent returned an error: %v\n", function_return);
			}
		} else{
			log.Printf("Failure: NewEvent returned an error: %v\n", function_return);
		}
		//Create event path_test
		function_return = NewEvent( "event_listener:path_test", map[string]interface{}{} );
		if( function_return.NoError() == true ){
			log.Printf("Success: Event event_lister:path_test created successfully.\n",);
			event = function_return.Data["event"].(Event_struct);
			function_return = event_dispatcher.PushEvent(event);
			if( function_return.NoError() == true ){
				log.Printf("Success: Event path_test pushed successfully.\n");
			} else{
				log.Printf("Failure: event_dispatcher.PushEvent returned an error: %v\n", function_return);
			}
		} else{
			log.Printf("Failure: NewEvent returned an error: %v\n", function_return);
		}
		//Create event regexpath_test
		function_return = NewEvent( "event_listener:regexpath_test", map[string]interface{}{} );
		if( function_return.NoError() == true ){
			log.Printf("Success: Event event_lister:regexpath_test created successfully.\n",);
			event = function_return.Data["event"].(Event_struct);
			function_return = event_dispatcher.PushEvent(event);
			if( function_return.NoError() == true ){
				log.Printf("Success: Event regexpath_test pushed successfully.\n");
			} else{
				log.Printf("Failure: event_dispatcher.PushEvent returned an error: %v\n", function_return);
			}
		} else{
			log.Printf("Failure: NewEvent returned an error: %v\n", function_return);
		}
		//Create event regex-test
		function_return = NewEvent( "event_listener:regex-test", map[string]interface{}{} );
		if( function_return.NoError() == true ){
			log.Printf("Success: Event event_lister:regex-test created successfully.\n",);
			event = function_return.Data["event"].(Event_struct);
			function_return = event_dispatcher.PushEvent(event);
			if( function_return.NoError() == true ){
				log.Printf("Success: Event regex-test pushed successfully.\n");
			} else{
				log.Printf("Failure: event_dispatcher.PushEvent returned an error: %v\n", function_return);
			}
		} else{
			log.Printf("Failure: NewEvent returned an error: %v\n", function_return);
		}
		//Create event 01_test
		function_return = NewEvent( "event_listener:01_test", map[string]interface{}{} );
		if( function_return.NoError() == true ){
			log.Printf("Success: Event event_lister:01_test created successfully.\n",);
			event = function_return.Data["event"].(Event_struct);
			function_return = event_dispatcher.PushEvent(event);
			if( function_return.NoError() == true ){
				log.Printf("Success: Event 01_test pushed successfully.\n");
			} else{
				log.Printf("Failure: event_dispatcher.PushEvent returned an error: %v\n", function_return);
			}
		} else{
			log.Printf("Failure: NewEvent returned an error: %v\n", function_return);
		}
	} else{
		t.Fail();
		log.Printf("Failure: NewEventDispatcher returned an error: %v \n", function_return);
	}

	//ProcessEvents
	function_return = event_dispatcher.ProcessEvents();
	if( function_return.NoError() == true ){
		log.Printf("Success: Events processed successfully\n");
	} else{
		t.Fail();
		log.Printf("Failure: event_dispatcher.ProcessEvents returned an error: %v\n", function_return);
		if( function_return.CodeEqual( ERROR_CODE_EVENT_PROCESSING_ERROR ) == true ){
			function_return = function_return.GetWrappedBottom();
			log.Printf("Failure: Wrapped error report: %v\n", function_return);
		}
	}

	//Return
}


//# Private Functions


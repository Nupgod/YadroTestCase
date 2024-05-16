package logger

import "sort"

func (c *Club) HandleEvents() {
	for _, event := range c.Events {
		c.Result = append(c.Result, event)
		switch event.ID {
			case 1:
				c.processArrival(event)
			case 2:
				c.processSit(event)
			case 3:
				c.processWait(event)
			case 4:
				c.processLeave(event)
		}
	}
	c.processEndOfDay()
}

func (c *Club) processArrival(event Event) {
	if event.Time.Before(c.StartTime) || event.Time.After(c.EndTime) {
		outEvent := Event{
			Time:    event.Time,
			ID:      13,
			Client:    "NotOpenYet",
		}
		c.Result = append(c.Result, outEvent)
		return
	}

	if c.Clients[event.Client] {
		outEvent := Event{
			Time:    event.Time,
			ID:      13,
			Client:    "YouShallNotPass",
		}
		c.Result = append(c.Result, outEvent)
		return
	}
	c.Clients[event.Client] = true
}

func (c *Club) processSit(event Event) {
	if !c.Clients[event.Client] {
		outEvent := Event{
			Time:    event.Time,
			ID:      13,
			Client:    "ClientUnknown",
		}
		c.Result = append(c.Result, outEvent)
		return
	}
	
	if  c.Tables[event.TableNum-1].Client != "" {
		outEvent := Event{
			Time:    event.Time,
			ID:      13,
			Client:    "PlaceIsBusy",
		}
		c.Result = append(c.Result, outEvent)
		return
	}
	for i := range c.Tables {
		if c.Tables[i].Client == event.Client {
			c.calculateRevenue(event, i)
			c.Tables[i].Client = ""
		}
	}
	c.Tables[event.TableNum-1].Client = event.Client
	c.Tables[event.TableNum-1].Start = event.Time
}

func (c *Club) processWait(event Event) { 
	if !c.Clients[event.Client] {
		outEvent := Event{
			Time:    event.Time,
			ID:      13,
			Client:    "ClientUnknown",
		}
		c.Result = append(c.Result, outEvent)
		return
	}
	c.Queue = append(c.Queue, event.Client)
	for _, table := range c.Tables {
		if table.Client == "" {
			outEvent := Event{
				Time:    event.Time,
				ID:      13,
				Client:    "ICanWaitNoLonger!",	
		}
		c.Queue = c.Queue[1:]
		c.Result = append(c.Result, outEvent)
		return
		}
	}
	if len(c.Queue) >= len(c.Tables) {
		outEvent := Event{
			Time:    event.Time,
			ID:      11,
			Client:    event.Client,
		}
		c.Result = append(c.Result, outEvent)
		delete(c.Clients, event.Client) 
		c.Queue = c.Queue[:len(c.Queue)-1]
	}
}

func (c *Club) processLeave(event Event) {
	if !c.Clients[event.Client] {
		outEvent := Event{
			Time:    event.Time,
			ID:      13,
			Client:    "ClientUnknown",
		}
		c.Result = append(c.Result, outEvent)
		return
	}
	for j := range c.Tables {
		if c.Tables[j].Client == event.Client {
			c.calculateRevenue(event, j)
			if len(c.Queue) > 0 {
				c.Tables[j].Client = c.Queue[0]
				c.Tables[j].Start = event.Time
				outEvent := Event{
					Time:    event.Time,
					ID:      12,
					Client:    c.Queue[0],
					TableNum: j+1,
				}
				c.Result = append(c.Result, outEvent)
				c.Queue = c.Queue[1:]
			} else { c.Tables[j].Client = ""} 
			break
		}
	}
	delete(c.Clients, event.Client)
}

func (c *Club) processEndOfDay() {
	var remainingClients []string
	for client := range c.Clients {
		remainingClients = append(remainingClients, client)
	}
	sort.Strings(remainingClients)
	for _, client := range remainingClients {
		for tableNum, table := range c.Tables {
			if table.Client == client {
				outEvent := Event{
					Time:   c.EndTime,
					ID:     11,
					Client: client,
				}
				c.calculateRevenue(outEvent, tableNum)
				c.Result = append(c.Result, outEvent)
				delete(c.Clients, client)
				break
			}
		}
		for i, waiter := range c.Queue {
			if client == waiter {
				c.Queue = append(c.Queue[:i], c.Queue[i+1])
				outEvent := Event{
					Time:   c.EndTime,
					ID:     11,
					Client: client,
				}
				c.Result = append(c.Result, outEvent)
			}
		}
		
	}
}
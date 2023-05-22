package examples

import "sync"

// Данный пример является копией предыдущего, но
// с горутинами и объяснением почему это плохая идея.
// Проблема заключается в том, что мы лочим доступ к мапе,
// и соответственно не можем проверять наличие других
// ключей в других горутинах - т.е. ничем не отличается
// от линейного выполнения.
// Но, если не ошибаюсь, можно блочить только изменение
// залоченной сущности, а чтение будет доступно. Но
// получится так, что при наличии большого числа повторений,
// проверки на наличие ключа смогут идти впереди внесения
// ключей - с точки зрения результата функции все будет
// корректно, но будут лишние операции с изменением ключа. т.е.
// мы будем заходить в if когда ключ существует, и тем самым
// занимать время программы фактом того что мы зашли в if
// и что-то там делаем, так и тем, что мы лочим сущность
// на добавление новых ключей.
func Set2(s []string) map[string]void {
	r := map[string]void{}

	var wg sync.WaitGroup
	wg.Add(len(s))

	var mux sync.Mutex

	for _, value := range s {
		go func(key string) {
			defer wg.Done()

			if _, ok := r[key]; !ok {
				mux.Lock()
				r[key] = v
				mux.Unlock()
			}
		}(value)

	}
	wg.Wait()
	return r
}